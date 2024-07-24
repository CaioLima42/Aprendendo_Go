package main

import (
	"log"
	"net"
	"time"
	"fmt"
)

const Port = "8080"
const SaveMode = true
const MessageRate = 1.0
const BanLimit = 10*60.0
const StrikeLimit = 10


func sensitive(message string) string{
	if SaveMode {
		return "[REDACTED]" 
	} else{
		return message
	}  
}

func saveRemoteAddr(conn net.Conn) string{
	if SaveMode {
		return "[REDACTED]" 
	} else{
		return conn.RemoteAddr().String()
	}  
}

type MessageType int
const(
	ClientConnected MessageType = iota + 1
	ClientDisconected
	NewMessage
)

type Message struct{
	Type MessageType
	Conn net.Conn
	Text string
}

type Client struct{
	Conn net.Conn
	LastMessage time.Time
	StrikeCount int
}

func server(messages chan Message){
	clients := map[string]*Client{}
	bannedMfs := map[string]time.Time{}
	for{
		msg := <- messages
		switch msg.Type{
		case ClientConnected:
			addr := msg.Conn.RemoteAddr().(*net.TCPAddr)
			bannedAt, banned := bannedMfs[addr.IP.String()]
			now := time.Now()
			if banned{
				if now.Sub(bannedAt) >= BanLimit {
					delete(bannedMfs, addr.IP.String())
					banned = false
				}
			}
			if !banned{
				log.Printf("Client %s connected", sensitive(addr.String()))
				clients[msg.Conn.RemoteAddr().String()] = &Client{
					Conn: msg.Conn,
					LastMessage: time.Now(),
				}
			} else {
				msg.Conn.Write([]byte(fmt.Sprintf("You are Banned :%f secs left\n", now.Sub(bannedAt).Seconds())))
				msg.Conn.Close()
			}
		case ClientDisconected:	
			addr := msg.Conn.RemoteAddr().(*net.TCPAddr)		
			log.Printf("Client %s disconnected", sensitive(addr.String()))
			delete(clients, addr.String())
		case NewMessage:
			now := time.Now()
			authorAddr := msg.Conn.RemoteAddr().(*net.TCPAddr)
			author := clients[authorAddr.String()]
			if author != nil{
			if now.Sub(author.LastMessage).Seconds() >= MessageRate{
				author.LastMessage = now
				author.StrikeCount = 0
				log.Printf("Client %s send message %s", sensitive(authorAddr.String()), msg.Text)
				for _, client := range clients{
					if client.Conn.RemoteAddr().String() != authorAddr.String(){
						_, err := client.Conn.Write([]byte(msg.Text))
						if err != nil{ 
							log.Printf("Could not send data to %s: %s\n", sensitive(client.Conn.RemoteAddr().String()), err)

						}
					}
				}
		} else{
			author.StrikeCount += 1 
			if author.StrikeCount >= StrikeLimit{
				bannedMfs[authorAddr.IP.String()] = now
				author.Conn.Write([]byte("You are banned\n"))
				author.Conn.Close()
			}
		}
	} else{
		msg.Conn.Close()
	}
}
}
}

func client(conn net.Conn, messages chan Message){
	buffer := make([]byte, 64)
	for{
		n, err := conn.Read(buffer)
		if err != nil{
			log.Printf("Could not read from a client %s: %s\n", saveRemoteAddr(conn), err)
			conn.Close()
			messages <- Message{
				Type: ClientDisconected,
				Conn: conn,
			}
			return
		}
		text := string(buffer[0:n])
	    messages <- Message{
			Type: NewMessage,
			Conn: conn,
			Text: text,
		}

	}

}

func main(){
	ln, err := net.Listen("tcp", ":" + Port)
    if err != nil {
        log.Fatalf("Error: could not listen a port %s: %s\n", Port, sensitive(err.Error()))
    }
	log.Printf("Listen to TCP connections on port %s...\n", Port)

	messages := make(chan Message)
	go server(messages)
    for {
        conn, err := ln.Accept()
        if err != nil {
            log.Printf("Could not accept a connection %s\n", sensitive(err.Error()))
        }
		log.Printf("Accepted connection from %s", sensitive(conn.RemoteAddr().String()))
		messages <- Message{
			Type: ClientConnected,
			Conn: conn,
		}
        go client(conn, messages)
    }
}