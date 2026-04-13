/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}

	if root.Val == subRoot.Val {
		val :=  isSame(root, subRoot)
		if val {
			return val
		}
	}

	

	left, right := root.Left, root.Right
	return isSubtree(left, subRoot) || isSubtree(right, subRoot)
}

func isSame(a1 *TreeNode, a2 *TreeNode) bool {
	if a1 == nil && a2 == nil {
		return true
	}

	if a1 == nil && a2 != nil {
		return false
	}

	if a1 != nil && a2 == nil {
		return false
	}

	if a1.Val != a2.Val {
		return false
	}

	q1 := isSame(a1.Left, a2.Left)
	q2 := isSame(a1.Right, a2.Right)

	return q1 && q2
}
