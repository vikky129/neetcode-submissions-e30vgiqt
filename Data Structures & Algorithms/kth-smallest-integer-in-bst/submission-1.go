/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
	_, g :=  PreOrder(root, 0, k)
	return g
}

func PreOrder(node *TreeNode, k, ck int) (int, int) {
	if node == nil {
		return k, -1
	}
	k, ff := PreOrder(node.Left, k, ck)
	if ff != -1 {
		return k, ff
	}
	k++
	if k == ck {
		return k, node.Val
	}
	k, ff = PreOrder(node.Right, k , ck)
	if ff != -1 {
		return k, ff
	}

	return k, -1
}
