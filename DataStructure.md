# Linked List :
Pointer to a node Struct is a Head Pointer.
In linked list we always need two node, one for storing the node and one for keeping track of its head, which we need to return at last.

**Ex - 
type Node struct {
    value int
    next  *Node
}

**head = *Node
list1 = *Node  ---> here list1 is also a head pointer to node struct****

  A linked list is a data structure that consists of a sequence of nodes, where each node contains a value and a reference to the next node in the sequence. The first node in the sequence is called the head of the linked list.
Linked lists are useful when the size of the data to be stored is not known in advance or when efficient insertions and deletions are required. The main advantage of linked lists over arrays is that the insertion and deletion operations can be done in constant time, O(1), whereas in arrays these operations are O(n) where n is the number of elements in the array.
There are several types of linked lists, including singly linked lists, doubly linked lists, and circular linked lists. In a singly linked list, each node has a reference to the next node in the sequence, while in a doubly linked list, each node has references to both the next and previous nodes in the sequence. A circular linked list is a linked list where the last node in the sequence points to the first node, forming a circle.
The following is an example of a singly linked list with three nodes:

**Example:**
**head -> [value: 1, next: node2] -> [value: 2, next: node3] -> [value: 3, next: nil]**

In this example, head is a reference to the first node in the list, which has a value of 1 and a reference to the second node. The second node has a value of 2 and a reference to the third node, and so on. The last node in the list has a value of 3 and a reference to nil, indicating the end of the list.

    Use slow pointer, fast pointer approch to resolve the linked list for detecting the loop, finding the mid element in linked list.
    Use Hash Table for solving the DS question.
    
**Tree Traversal Method:**
Pre-order : root --> left --> right,
In-order : left --> root --> right,
Post-order : left --> right --> root

**Binary Tree**: A binary tree is a non linear data structure, 
A binary tree is a type of tree data structure in which each node can have at most two children, commonly referred to as the left child and the right child. The binary tree follows a hierarchical structure, with a single node at the top called the root node. Each node in the binary tree can have zero, one, or two children.

Here are a few key characteristics of a binary tree:

1- Root Node: The topmost node in the binary tree is called the root node. It serves as the starting point for traversing the tree.
2- Parent and Child Nodes: Each node in a binary tree can have at most two children. These children are referred to as the left child and the right child. A node that has a child is considered the parent of that child.
3- Leaf Nodes: Leaf nodes are nodes that have no children. They are located at the bottom of the tree.
4- Internal Nodes: Internal nodes are nodes that have at least one child. They are located between the root node and the leaf nodes.

Binary trees are commonly used to represent hierarchical structures, such as binary search trees, expression trees, and decision trees. They provide efficient searching, insertion, and deletion operations when appropriately balanced.

Example: 
       4
     /   \
    2     6
   / \   / \
  1   3 5   7

****Binary Search Tree** - use binary serarch to get o(logn), time complexity. it divide the serarch in half and do the iterarion.
if the list is sorted always try to use binary serarhc on them.
Binary search tree is a binary tree in which a node have a value greater than all values in its left subtree and smaller than all values in its right subtree
In a binary search tree, each node has **at most two children**: a left child and a right child. The left child of a node contains a value that is less than the value of the node, and the right child contains a value that is greater than the value of the node.

In the case of a single-element binary search tree [0], there is only one node, and it does not have any children. Since there are no children to compare values with, the condition for a binary search tree is automatically satisfied, and thus, it is a valid binary search tree.

**Assuming an empty tree is considered a valid BST.**

Use recursion and memorization(using cache) to resolve the dynamic programming issue.for fabonanci and other issue.

