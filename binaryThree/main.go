package main

import (
	"fmt"
)

type Node struct{
	data string
	left *Node
	right *Node
}
func InorderPrint(root *Node) {
	if root == nil {
	 return
	}
	if root.left != nil {
		InorderPrint(root.left)
	}
	fmt.Print(root.data, " ")
	if root.right != nil {
		InorderPrint(root.right)
	}
}

func PreorderPrint(root *Node) {
	if root == nil {
	 return
	}
	fmt.Print(root.data, " ")
	if root.left != nil {
		PreorderPrint(root.left)
	}
	if root.right != nil {
		PreorderPrint(root.right)
	}
}

func Postorder(root *Node) {
	if root == nil {
	 return
	}
	if root.left != nil {
		Postorder(root.left)
	}
	if root.right != nil {
		Postorder(root.right)
	}
	fmt.Print(root.data, " ")
}
func main(){
	A := Node{data: "A"}
	B := Node{data: "B"}
	C := Node{data: "C"}
	D := Node{data: "D"}
	E := Node{data: "E"}
	F := Node{data: "F"}
	G := Node{data: "G"}

	A.left = &B
	A.right = &C
	B.left = &D
	B.right = &E
	C.left = &F
	C.right = &G
	fmt.Println("Inorder: ")
	InorderPrint(&A)
	fmt.Println()
	fmt.Println("Preorder: ")
	PreorderPrint(&A)
	fmt.Println()
	fmt.Println("Posorder: ")
	Postorder(&A)

}