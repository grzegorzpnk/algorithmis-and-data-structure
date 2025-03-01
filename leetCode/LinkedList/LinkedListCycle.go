package LinkedList

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {

	if head == nil {
		return false
	}

	rememberedNodesMap := make(map[*ListNode]bool)
	current := head

	for current.Next != nil {
		if rememberedNodesMap[current] {
			return true
		}
		rememberedNodesMap[current] = true
		current = current.Next
	}

	return false
}
