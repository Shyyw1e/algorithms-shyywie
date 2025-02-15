package exercises

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type SingleLinkedList struct {
	head *Node
}

func (list *SingleLinkedList) Append(value int) {
	newNode := &Node{data : value}
	if list.head == nil {
		list.head = newNode
		return
	}
	current := list.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode

}

func (list *SingleLinkedList) Delete(value int) {
	if list.head == nil {
		return
	}
	if list.head.data == value {
		list.head = list.head.next
		return 
	}
	current := list.head
	if current.next != nil && current.next.data != value {
		current = current.next
	}
	if current.next != nil {
		current.next = current.next.next
	}
}