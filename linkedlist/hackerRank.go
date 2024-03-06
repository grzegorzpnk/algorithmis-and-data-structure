package linkedlist

//
//func reversePrint(llist *LinkedList) {
//
//	var stack []int32
//
//	for llist != nil {
//		stack = append(stack, llist.data)
//		llist = llist.next
//	}
//
//	for len(stack) > 0 {
//		fmt.Println(stack[len(stack)-1])
//		stack = stack[:len(stack)-1]
//	}
//
//}

//func reverserDLL(dll *DoubledLinkedList) {
//
//	var stack []int32
//
//	for dll != nil {
//		stack = append(stack, dll)
//		llist = llist.next
//	}
//
//	for len(stack) > 0 {
//
//		dll.
//			stack = stack[:len(stack)-1]
//	}
//
//}
//

//func reverseDLL(Head *NodeDoubled) *NodeDoubled {
//	var prev, next *NodeDoubled
//	current := Head
//
//	for current != nil {
//		next = current.next
//
//
//	}
//
//	for current != nil {
//		next = current.next
//		current.next = prev
//		current.prev = next
//		prev = current
//		current = next
//	}
//
//	return prev
//}

func ReverseDLL(head *NodeDoubled) *NodeDoubled {
	var temp *NodeDoubled
	current := head

	for current != nil {
		temp = current.next
		current.next = current.prev
		current.prev = temp

		current = current.prev
	}

	return current
}

func (dll *DoubledLinkedList) Reverse() {

	current := dll.Head
	var temp *NodeDoubled

	for current != nil {
		temp = current.next
		current.next = current.prev
		current.prev = temp

		current = current.prev
	}

	temp = dll.Head
	dll.Head = dll.Tail
	dll.Tail = temp

}
