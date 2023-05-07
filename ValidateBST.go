// A valid BST is defined as follows:
// The left subtree of a node contains only nodes with keys less than the node's key.
// The right subtree of a node contains only nodes with keys greater than the node's key.
// Both the left and right subtrees must also be binary search trees.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return validate(root, nil, nil)
}

func validate(node *TreeNode, min *TreeNode, max *TreeNode) bool {
	if node == nil {
		return true
	}

	if (min != nil && node.Val <= min.Val) || (max != nil && node.Val >= max.Val) {
		return false
	}

	return validate(node.Left, min, node) && validate(node.Right, node, max)
}
