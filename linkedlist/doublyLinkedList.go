package linkedlist

import "fmt"

type NodeDoubled struct {
	value int
	next  *NodeDoubled
	prev  *NodeDoubled
}

type DoubledLinkedList struct {
	head *NodeDoubled
	tail *NodeDoubled
}

//insert new element at the end of the existing list

func (dll *DoubledLinkedList) insertFront(data int) {
	newNode := &NodeDoubled{value: data, next: nil, prev: nil}

	if dll.head == nil {
		dll.head = newNode
		dll.tail = newNode
	} else {
		newNode.next = dll.head
		dll.head.prev = newNode
		dll.head = newNode
	}
}

func (dll *DoubledLinkedList) insertEnd(data int) {
	newNode := &NodeDoubled{value: data, next: nil, prev: nil}

	if dll.head == nil {
		dll.head = newNode
		dll.tail = newNode
	} else {
		newNode.prev = dll.tail
		dll.tail.next = newNode
		dll.tail = newNode
	}
}

func (dll *DoubledLinkedList) insertAt(data, position int) {
	newNode := &NodeDoubled{value: data, next: nil, prev: nil}

	if position <= 0 {
		dll.insertFront(data)
		return
	}

	current := dll.head
	for i := 1; i < position && current != nil; i++ {
		current = current.next
	}

	if current == nil {
		dll.insertEnd(data)
		return
	}

	newNode.next = current.next
	newNode.prev = current
	current.next = newNode

	if newNode.next != nil {
		newNode.next.prev = newNode
	}

}

func (dll *DoubledLinkedList) delete(data int) {

	current := dll.head

	for current != nil && current.value != data {
		current = current.next
	}

	if current == nil {
		fmt.Printf("Not found such a Node")
	}

	//to distinguish if the deleted node is the head
	if current.prev != nil {
		current.prev.next = current.next
	} else {
		dll.head = current.next
	}

	//to distinguish if the deleted node is the tail
	if current.next != nil {
		current.next.prev = current.prev
	} else {
		dll.tail = current.prev
	}

	if dll.head == nil {
		dll.tail = nil
	}

}

func (dll *DoubledLinkedList) search(value int) *NodeDoubled {

	current := dll.head

	for current != nil && current.value != value {
		current = current.next
	}

	if current == nil {
		fmt.Errorf("node not found")
	}
	return current
}

func (dll *DoubledLinkedList) display() {

	current := dll.head
	for current.next != nil {
		fmt.Printf("%d -> ", current.value)
		current = current.next
	}
	fmt.Println("finish")
}
