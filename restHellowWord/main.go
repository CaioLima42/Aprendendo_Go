package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Print("Hello world received a request.")
	log.Print(r.Header["Sec-Ch-Ua"])
	defer log.Print("End hello world request")
	fmt.Fprintf(w, "Flamengo")
}

func main() {
	log.Print("Hello world sample started.")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

