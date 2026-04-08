/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var start *ListNode
	var curr *ListNode

	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	var carry int
	var sum int

	for l1 != nil && l2 != nil {
		sum = l1.Val + l2.Val  + carry
		if sum >= 10 {
			sum = sum%10
			carry = 1
		} else {
			carry = 0
		}

		if start == nil {
			start = &ListNode {
				Val: sum, 
				Next: nil,
			}
			curr = start
			l1 = l1.Next
			l2 = l2.Next
			continue
		}

		curr.Next = &ListNode {
				Val: sum, 
				Next: nil,
			}

		curr = curr.Next
		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		sum = l1.Val  + carry
		if sum >= 10 {
			sum = sum%10
			carry = 1
		} else {
			carry = 0
		}

		curr.Next = &ListNode {
				Val: sum, 
				Next: nil,
			}

		curr = curr.Next
		l1 = l1.Next
	}

	for l2 != nil {
		sum = l2.Val  + carry
		if sum >= 10 {
			sum = sum%10
			carry = 1
		} else {
			carry = 0
		}

		curr.Next = &ListNode {
				Val: sum, 
				Next: nil,
			}

		curr = curr.Next
		l2 = l2.Next
	}

	if carry > 0 {
		curr.Next = &ListNode {
				Val: carry, 
				Next: nil,
			}
	}

	return start

}
