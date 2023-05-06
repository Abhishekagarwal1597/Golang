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

    var result [][]int
    queue := []*TreeNode{root}
    

    for len(queue) > 0 {
        levelSize := len(queue)
        levelValues := make([]int, levelSize)
        fmt.Println("queueLength:", len(queue) )

        for i := 0; i < levelSize; i++ {
            node := queue[i]
            levelValues[i] = node.Val

            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }

        result = append(result, levelValues)
        queue = queue[levelSize:]
    }

    return result
}

### In this code, we define the TreeNode struct to represent a node in the binary tree. Each node has a value (Val) and pointers to its left and right children (Left and Right).
The levelOrder function takes the root node of the binary tree as input and returns a 2D slice representing the level order traversal of the tree.
We start by checking if the root node is nil, in which case we return an empty 2D slice. Otherwise, we initialize an empty result slice and a queue to hold the nodes.
We perform a standard breadth-first search (BFS) traversal using a loop. In each iteration of the loop, we determine the number of nodes in the current level (levelSize). We create a temporary slice levelValues to store the values of the nodes at the current level.
We iterate through the nodes in the current level, adding their values to levelValues and enqueueing their children if they exist.
After processing all the nodes in the current level, we append levelValues to the result slice and update the queue by removing the nodes of the current level.
Finally, we return the result slice containing the level order traversal of the binary tree.
In the main function, we create an example binary tree and call the levelOrder function to perform the level order traversal. The result is then printed to the console.
