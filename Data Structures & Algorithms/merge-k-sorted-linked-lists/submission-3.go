/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0  {
		return nil
	}

	if len(lists) == 1 {
		return lists[0]
	}

	var start = lists[0]
	var currNode *ListNode
	for i := 1; i < len(lists); i++ {
		currNode = lists[i]
		start = MergeTwoLists(start, currNode)
	}

	return start
}

func MergeTwoLists(s1, s2 *ListNode) *ListNode{
	var newStart *ListNode
	var curr *ListNode
	for s1 != nil && s2 != nil {
		if s1.Val <= s2.Val {
			if newStart == nil {
				newStart = s1
				curr = s1
				s1 = s1.Next
				continue
			}

			curr.Next = s1
			curr = curr.Next
			s1 = s1.Next
		} else {
			if newStart == nil {
				newStart = s2
				curr = s2
				s2 = s2.Next
				continue
			}

			curr.Next = s2
			curr = curr.Next
			s2 = s2.Next
		}
	}

	for s1 != nil {
		if newStart == nil {
			newStart = s1
			curr = s1
			s1 = s1.Next
			continue
		}

		curr.Next = s1
		curr = curr.Next
		s1= s1.Next
	}

	for s2 != nil {
		if newStart == nil {
			newStart = s2
			curr = s2
			s2 = s2.Next
			continue
		}

		curr.Next = s2
		curr = curr.Next
		s2 = s2.Next
	}

	return newStart
}
