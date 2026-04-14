/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {

	if root == nil {
		return []int{}
	}
    li := []int{}
	mp := map[int]struct{}{}

	li = append(li, root.Val)
	mp[0] = struct{}{}

	return print(root, li, mp, 0)
}

func print(root *TreeNode, li []int, mp map[int]struct{}, level int) []int {
	if root == nil {
		return li
	}
	if _, ok := mp[level]; !ok {
		li = append(li, root.Val)
		mp[level] = struct{}{}
	}

	li = print(root.Right, li, mp, level + 1)
	li = print(root.Left, li, mp, level + 1)

	return li
}
