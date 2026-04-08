/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
    var sp, mp = head, head

	for {
		if sp == nil || sp.Next == nil || mp == nil || mp.Next == nil  || mp.Next.Next == nil{
			return false
		}

		sp = sp.Next
		mp = mp.Next.Next

		if sp == mp {
			return true
		}
	}
	return false
}
