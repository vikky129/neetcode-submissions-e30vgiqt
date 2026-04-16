/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
	var max = -1000

	return goodCount(root, max, 0)

}

func goodCount(node *TreeNode, max int, count int) int {
	if node == nil {
		return count
	}

	if node.Val >= max {
		count++
		max = node.Val
	}

	count1 := goodCount(node.Left, max, count)
	count2 := goodCount(node.Right, max, count)

	return count1 + count2 - count
}
