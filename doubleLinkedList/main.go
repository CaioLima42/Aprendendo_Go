package main

import (
	"fmt"
)

type Node struct{
	value int
	Next *Node
	Previos *Node 
}
func Apeend(principal *Node, newNode * Node){
	root := principal
	for (root.Next != nil){
		root = root.Next
	}
	root.Next = newNode
	}
func Print(principal *Node){
	root := principal
	for root.Next != nil{
		fmt.Println(root.value)
		root = root.Next
	}
   	fmt.Println(root.value)
}
func Search(principal *Node, loc int) *Node{
	root := principal
	cont := 0
	for root.Next != nil{
		if loc == cont{
			return root
		}
	}
	return nil
}
func Len(principal * Node) int{
	root := principal
	cont := 0
	for root.Next != nil{
		cont++
		root = root.Next
	}
	cont++
	return cont 
}
func Insert(root *Node, newNode *Node, loc int) *Node {
	counter := 0
	n := root
	var p *Node
	for n != nil {
		if counter == loc {
			newNode.Next = n
			newNode.Previos = p
		if n.Next != nil {
			n.Next.Previos = newNode
		}
		if p != nil {
			p.Next = newNode
			return root
		}
		return newNode
		}
		p = n
		n = n.Next
		counter = counter + 1
		}
		if counter == loc {
			newNode.Previos = p
			p.Next = newNode
			return root
		}
	return nil
}
func Delete(root *Node, loc int) *Node {
	counter := 0
	n := root
	var p *Node
	for n != nil {
		if counter == loc {
			if p == nil {
			temp := n.Next
			n.Next = nil
			n.Previos = nil
			return temp
			}
		p.Next = n.Next
		n.Next = nil
		n.Previos = nil
		return root
		}
	p = n
	n = n.Next
	counter = counter + 1
	}
	return nil
}

func main(){
	a := Node{value: 1}
	b := Node{value: 2}
	c := Node{value: 3}
	
	a.Next = &b
	b.Next = &c
	b.Previos = &a
	c.Previos = &b 
	newNode := Node{value: 4}
	Print(&a)
	Insert(&a, &newNode, 1)
	Print(&a)

}