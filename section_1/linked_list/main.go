package main

import (
	"fmt"
)


type List struct{
	head *Node
}

func(l *List) First() *Node{
	return l.head
}

func(l *List) Push(value int){

}

type Node struct{
	value int
	next *Node
}

func(n *Node) Next(params)