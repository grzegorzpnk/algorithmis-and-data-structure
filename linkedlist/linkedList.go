package linkedlist

import "fmt"

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
}

//insert new element at the end of the existing list

func (ll *LinkedList) insert(data int) {

	newNode := &Node{value: data, next: nil}

	if ll.head == nil {
		ll.head = newNode
	} else {
		current := ll.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
}

func (ll *LinkedList) delete(key int) {

	if ll.head == nil {
		return
	}

	if ll.head.value == key {
		ll.head = ll.head.next
		return
	}

	prev := ll.head
	current := ll.head.next

	for current != nil {
		if current.value == key {
			prev.next = current.next
			return
		}
		prev = current
		current = current.next
	}
}

func (ll *LinkedList) search(value int) *Node {

	current := ll.head

	for current != nil {
		if current.value == value {
			return current
		}
		current = current.next
	}

	return nil
}

func (ll *LinkedList) display() {

	current := ll.head
	for current.next != nil {
		fmt.Printf("%d -> ", current.value)
		current = current.next
	}
	fmt.Println("finish")
}
