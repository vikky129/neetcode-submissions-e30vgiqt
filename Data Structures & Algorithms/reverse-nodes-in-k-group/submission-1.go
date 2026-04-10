/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 1, 2, 3, 4, 5, 6
func reverseKGroup(head *ListNode, k int) *ListNode {
    if k == 1 {
		return head
	}
	
	var kStart *ListNode = head
	var newStart *ListNode
	var curr = 1
	var nStart *ListNode
	var prevNStart *ListNode
	for head != nil {
		head = head.Next
		curr++
		if head == nil {
			break
		}
		if curr == k {
			if newStart == nil {
				newStart = head
			}
			var ed = head.Next
			var hd1 = head
			fmt.Println(kStart.Val, head.Val)
			nStart = reverse(kStart, head)
			if prevNStart != nil {
				prevNStart.Next = hd1
			}
			nStart.Next = ed
			head = ed
			kStart = head
			curr = 1
			prevNStart = nStart
		}
	}

	return newStart

}



// 1, 2, 3, 4, 5
func reverse(start, end *ListNode) *ListNode{
	if start == end {
		return start
	}

	endNode := reverse(start.Next,  end)
	endNode.Next = start
	return start
}