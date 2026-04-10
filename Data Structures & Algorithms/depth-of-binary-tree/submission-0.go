/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
    return depth(root, 1)
}


func depth(root *TreeNode, dep int) int {
	if root == nil {
		return dep - 1
	}

	d1 := depth(root.Left, dep + 1)
	d2 := depth(root.Right, dep + 1)

	if d1 >= d2 {
		return d1
	} else {
		return d2
	}
}
