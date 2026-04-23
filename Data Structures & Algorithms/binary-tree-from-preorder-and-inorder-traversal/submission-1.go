/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func buildTree(preorder []int, inorder []int) *TreeNode {
    var mp = make(map[int]int)

	node := &TreeNode {
		Val:preorder[0],
	}

	for i := 0; i < len(inorder); i++ {
		mp[inorder[i]] = i
	}

	var visited = make(map[int]struct{})
	visited[preorder[0]] = struct{}{}

	Mama(node, 1, mp, preorder, visited)
	return node
}



func Mama(node *TreeNode, i int,mp map[int]int, preorder []int, visited map[int]struct{}) int {

	if i == len(preorder) {
		return i
	}

	if node == nil {
		return i
	}

	if mp[preorder[i]] < mp[node.Val] {
		node.Left = &TreeNode{
			Val:preorder[i],
		}
		visited[preorder[i]] = struct{}{}
		i = Mama(node.Left, i + 1, mp, preorder, visited)
	}

	if i == len(preorder) {
		return i
	}

	if mp[preorder[i]] > mp[node.Val] {
		km := NearNode(visited, mp, preorder[i])
		if km == node.Val {
			if _, ok := visited[preorder[i]]; !ok {
				node.Right = &TreeNode{
					Val: preorder[i],
				}
				visited[preorder[i]] = struct{}{}

				i = Mama(node.Right, i+1, mp, preorder, visited)
			}
		}
	}

	return i
}

func NearNode(visited map[int]struct{}, mp map[int]int, cValue int) int {
	//var node *TreeNode
	var min int = len(mp)
	var k int
	for key, _ := range visited {
		mk := mp[cValue] - mp[key]
		if mk > 0 && mk < min {
			min = mk
			k= key
		}
	}

	return k
}