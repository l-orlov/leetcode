package has_path_sum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasSum(root *TreeNode, targetSum int) bool {
	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}

	var l, r bool
	if root.Left != nil {
		l = hasPathSum(root.Left, targetSum-root.Val)
	}
	if root.Right != nil {
		r = hasPathSum(root.Right, targetSum-root.Val)
	}

	return l || r
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	return hasSum(root, targetSum)
}
