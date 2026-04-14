/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
    if root == nil {
		return [][]int{}
	}

	var arr = [][]int{}

	if root.Left == nil && root.Right == nil {
		return [][]int{
			[]int{root.Val},
		}
	}

	nodes := []*TreeNode {root}
	//var cNodes []*TreeNode
	for {
		arr, nodes = Kaka(nodes, arr)
		if len(nodes) == 0 {
			break
		}
	}

	return arr
}

func Kaka(nodes []*TreeNode, result [][]int) ([][]int, []*TreeNode) {
	var t []int
	var k []*TreeNode
	for _, node := range nodes {
		t = append(t, node.Val)
		if node.Left != nil {
			k = append(k, node.Left)
		}

		if node.Right != nil {
			k = append(k, node.Right)
		}
	}

	result = append(result, t)
	return result, k
}
