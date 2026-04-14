/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
    return ancestor(root, p, q)
}

func ancestor(root, p, q *TreeNode) *TreeNode {
	if p.Val < root.Val && q.Val > root.Val {
		return root
	}

	if p.Val  > root.Val && q.Val < root.Val {
		return root
	}

	if p.Val == root.Val {
		return p
	}

	if q.Val == root.Val {
		return q
	}

	if root == nil {
		return nil
	}

	if root.Left != nil {
		left := ancestor(root.Left, p, q)
		if left != nil {
			return left
		}
	}

	if root.Right != nil {
		right := ancestor(root.Right, p, q)
		if right != nil {
			return right
		}
	}
	

	return nil
}
