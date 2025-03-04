package LinkedList

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	//new list
	newHead := &ListNode{}
	//remember value of head ptr
	current := newHead

	for list1 != nil || list2 != nil {
		//if-statement to take a node from list1
		if (list1 != nil && list2 != nil && list1.Val <= list2.Val) || list2 == nil {
			current.Next = list1
			list1 = list1.Next

			//if-statement to take a node from list1
		} else if list2 != nil {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	return newHead.Next

}
