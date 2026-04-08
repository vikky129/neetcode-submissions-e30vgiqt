/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next == nil {
		head = nil
		return head
	}

	_, head = DelN(head, nil, head, n)
	fmt.Println("main", head.Val)
	return head
}

func DelN(start *ListNode, prevNode *ListNode, head *ListNode, n int) (int, *ListNode) { 
	if start.Next == nil {
		if n == 1 {
			prevNode.Next = nil
		}
		return 1, head
	}

	r, head := DelN(start.Next, start, head, n)
	r++
	if r == n {	
		if prevNode == nil {
			head = head.Next
		} else {
			prevNode.Next = start.Next
		}
	}
	return r, head
}
