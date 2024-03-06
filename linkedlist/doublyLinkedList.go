package linkedlist

import "fmt"

type NodeDoubled struct {
	value int32
	next  *NodeDoubled
	prev  *NodeDoubled
}

type DoubledLinkedList struct {
	Head *NodeDoubled
	Tail *NodeDoubled
}

//insert new element at the end of the existing list

func (dll *DoubledLinkedList) InsertFront(data int32) {
	newNode := &NodeDoubled{value: data, next: nil, prev: nil}

	if dll.Head == nil {
		dll.Head = newNode
		dll.Tail = newNode
	} else {
		newNode.next = dll.Head
		dll.Head.prev = newNode
		dll.Head = newNode
	}
}

func (dll *DoubledLinkedList) InsertEnd(data int32) {
	newNode := &NodeDoubled{value: data, next: nil, prev: nil}

	if dll.Head == nil {
		dll.Head = newNode
		dll.Tail = newNode
	} else {
		newNode.prev = dll.Tail
		dll.Tail.next = newNode
		dll.Tail = newNode
	}
}

func (dll *DoubledLinkedList) InsertAt(data int32, position int) {
	newNode := &NodeDoubled{value: data, next: nil, prev: nil}

	if position <= 0 {
		dll.InsertFront(data)
		return
	}

	current := dll.Head
	for i := 1; i < position && current != nil; i++ {
		current = current.next
	}

	if current == nil {
		dll.InsertEnd(data)
		return
	}

	newNode.next = current.next
	newNode.prev = current
	current.next = newNode

	if newNode.next != nil {
		newNode.next.prev = newNode
	}

}

func (dll *DoubledLinkedList) Delete(data int32) {

	current := dll.Head

	for current != nil && current.value != data {
		current = current.next
	}

	if current == nil {
		fmt.Printf("Not found such a Node")
	}

	//to distinguish if the deleted node is the Head
	if current.prev != nil {
		current.prev.next = current.next
	} else {
		dll.Head = current.next
	}

	//to distinguish if the deleted node is the Tail
	if current.next != nil {
		current.next.prev = current.prev
	} else {
		dll.Tail = current.prev
	}

	if dll.Head == nil {
		dll.Tail = nil
	}

}

func (dll *DoubledLinkedList) Search(value int32) *NodeDoubled {

	current := dll.Head

	for current != nil && current.value != value {
		current = current.next
	}

	if current == nil {
		fmt.Errorf("node not found")
	}
	return current
}

func (dll *DoubledLinkedList) Display() {

	current := dll.Head
	for current != nil {
		fmt.Printf("%d -> ", current.value)
		current = current.next
	}
	fmt.Println("finish")
}
