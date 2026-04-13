/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
	_, isB := HeightDiff(root)

	return isB
}


func HeightDiff(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}

	dL, isNotBalL := HeightDiff(root.Left)
	dR, isNotBalR := HeightDiff(root.Right)

	if !isNotBalL || !isNotBalR {
		return -1, false
	}

	if mod(dR - dL) > 1 {
		return -1, false
	}

	if dL > dR {
		return dL + 1, true
	}

	return dR + 1, true
}



func mod(val int) int {
	if val < 0 {
		return val * -1
	}

	return val
}
