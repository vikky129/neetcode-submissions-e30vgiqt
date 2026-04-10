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

	var resultList = []*ListNode{}
	for  {
		if len(lists) == 1 {
			return lists[0]
		}

		resultList = []*ListNode{}
		for i := 0; i <= len(lists); i = i + 2 {
			if i >= len(lists) -1 {
				if i == len(lists) - 1 {
					resultList = append(resultList, lists[i])
				}
				break
			}
			fmt.Println(i, i+1, len(lists))
			resultList = append(resultList, MergeTwoLists(lists[i], lists[i+1]))
		}

		lists = resultList
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
