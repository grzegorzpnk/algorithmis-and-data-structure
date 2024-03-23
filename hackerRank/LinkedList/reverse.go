package LinkedList

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}

	var previousNode *ListNode = nil
	currentNode := head

	for currentNode != nil {
		//tmp := currentNode.Next
		//currentNode.Next = previousNode
		//previousNode = currentNode
		//currentNode = tmp

		tmp := currentNode.Next
		currentNode = previousNode
		previousNode = currentNode
	}

	return previousNode
}

//
//func reverseList(head *ListNode) *ListNode {
//	var previousNode *ListNode = nil
//	currentNode := head
//
//	for currentNode != nil {
//		nextNode := currentNode.Next
//		currentNode.Next = previousNode
//		previousNode = currentNode
//		currentNode = nextNode
//	}
//
//	return previousNode
//}
