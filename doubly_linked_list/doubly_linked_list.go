package main

import (
	"fmt"
)

type Node struct {
	data interface{}
	next *Node
	prev *Node
}

type DoublyLinkedList struct {
	first *Node
	last *Node
}

func (dll *DoublyLinkedList) InsertAtEnd(data interface{}) {
	newNode := &Node{data, nil, nil}

	if dll.first == nil {
		dll.first = newNode
		dll.last = newNode

	  return
	}

	dll.last.next = newNode
	newNode.prev = dll.last
	dll.last = newNode
}

func (dll *DoublyLinkedList) RemoveFromFront() interface{} {
	removed := dll.first
	dll.first = dll.first.next
	dll.first.prev = nil

	return removed.data
}

func (dll DoublyLinkedList) PrintEndToFront() {
	for currNode := dll.last; currNode != nil; currNode = currNode.prev {
		fmt.Println(currNode.data)
	}
}

func main() {
	list := &DoublyLinkedList{nil, nil}
	list.InsertAtEnd("foo")
	list.InsertAtEnd("fizz")
	list.InsertAtEnd("bar")
	list.InsertAtEnd("bazz")

	list.PrintEndToFront()
}
