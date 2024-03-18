package is_same_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}

	if p.Val != q.Val {
		return false
	}

	if p.Left == nil || q.Left == nil {
		if p.Left != q.Left {
			return false
		}
	}

	if p.Right == nil || q.Right == nil {
		if p.Right != q.Right {
			return false
		}
	}

	if p.Left != nil {
		if !isSameTree(p.Left, q.Left) {
			return false
		}
	}

	if p.Right != nil {
		if !isSameTree(p.Right, q.Right) {
			return false
		}
	}

	return true
}
