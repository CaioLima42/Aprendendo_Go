package main

import (
	"fmt"
	//"go/printer"
)

type Node struct{
	value int
	next *Node
}

func Print(principal *Node){
	root := principal
	for (root.next != nil){
		fmt.Println(root.value)
		root = root.next
	}
	fmt.Println(root.value)
	}
func Len(principal *Node){
	root := principal
	tamanho := 0
	for (root.next != nil){
		tamanho++
		root = root.next
	}
	tamanho++
	fmt.Println(tamanho)
	}
func Apeend(principal *Node, newNode * Node){
	root := principal
	for (root.next != nil){
		root = root.next
	}
	root.next = newNode
	}

func InsertStart(principal *Node, newNode * Node){
	aux := principal
	principal = newNode
	principal.next = aux
	}

func Delete(principal *Node, loc int) error {
	root := principal
	cont := 0
	for (root != nil){
		if (cont == loc-1){
			root.next = root.next.next
			return nil
		}
		cont++
	}
	return fmt.Errorf("went past no of elements in list")
}


func Insert(principal *Node, loc int, newNode *Node) error {
	root := principal
	cont := 0
	for (root.next != nil){
		if (loc-1 == cont){
			var aux Node = *root.next
			root.next = newNode
			newNode.next = &aux
			return nil
		}
		root = root.next
		cont++
	}
	return fmt.Errorf("went past no of elements in list")		
	}
func Search(primeiro *Node, val int) *Node {
	root := primeiro
	for primeiro != nil{
		if root.value == val{
			return root
		}
		root = root.next
	}
	return nil
}


func main(){
	primeiro := Node{value: 1}
	segundo := Node{value: 2}
	terceiro := Node{value: 3}
	primeiro.next = &segundo
	segundo.next = &terceiro
	//newNode := Node{value: 4} 
	//Apeend(&primeiro,&newNode)
	//Print(&primeiro)
	//InsertStart(&primeiro, &newNode) esta com erro
	fmt.Println(*Search(&primeiro, 1))
	Delete(&primeiro, 1)
	Print(&primeiro)
	fmt.Print(Search(&primeiro,3))
}