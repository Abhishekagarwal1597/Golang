Preorder - root->left->right

// Node represents a node in an n-ary tree with an integer value Val 
// and a slice of child nodes Children.
type Node struct {
    Val      int
    Children []*Node
}

// preorder returns a list of node values in pre-order traversal order for an n-ary tree 
// rooted at the input node root.
// If the input node root is nil, an empty list is returned.
func preorder(root *Node) []int {
    var result []int
    if root == nil {
        return result
    }
    // Append the value of the root node to the pre-order list.
    result = append(result, root.Val)
    // Traverse the subtree rooted at each child node in pre-order and append the 
    // resulting list to the pre-order list.
    for _, child := range root.Children {
        result = append(result, preorder(child)...)
    }
    // Return the pre-order list.
    return result
}

