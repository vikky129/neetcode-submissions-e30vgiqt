/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var start *ListNode
	return Rev(head, nil, start)
}



// 1,2,3,4,5
func Rev(head *ListNode, prevNode, start *ListNode) *ListNode {
	if head.Next == nil {
		start = head
		start.Next = prevNode
		return start
	}

	start = Rev(head.Next, head, start)
	head.Next = prevNode
	return start
}
