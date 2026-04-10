/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
    _, de := maxDepth(root)
	return de
}

func maxDepth(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}

	d1, maxDp1 := maxDepth(root.Left)
	d2, maxdp2 := maxDepth(root.Right)
	maxL := max(maxDp1, maxdp2, d1+ d2)

	if d1 >= d2 {
		return d1 + 1, maxL
	} else {
		return d2 + 1, maxL
	}
}


func max(a, b, c int) int {
	var tMax  int

	if a > b {
		tMax = a
	} else {
		tMax = b
	}

	if c > tMax {
		tMax = c
	} 

	return tMax
}