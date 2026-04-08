/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	var newStart *ListNode
	var runn *ListNode
    for list1 != nil && list2 != nil{
		if list1.Val <= list2.Val {
			if newStart == nil {
				newStart = list1
				runn = newStart
				list1 = list1.Next
				continue
			}
			runn.Next = list1
			list1 = list1.Next
			runn = runn.Next
		} else {
			if newStart == nil {
				newStart = list2
				runn = newStart
				list2 = list2.Next
				continue
			}
			runn.Next = list2
			list2 = list2.Next
			runn = runn.Next
		}
	}

	for list1 != nil {
		if newStart == nil {
				newStart = list1
				runn = newStart
				list1 = list1.Next
				continue
			}
		runn.Next = list1
		list1 = list1.Next
		runn = runn.Next
	}

	for list2 != nil {
		if newStart == nil {
				newStart = list2
				runn = newStart
				list2 = list2.Next
				continue
			}
		runn.Next = list2
		list2 = list2.Next
		runn = runn.Next
	}

	return newStart
}
