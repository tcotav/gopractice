package bst

import (
	"fmt"
)

/* Define the base Node struct */
type Node struct {
	value int
	left  *Node
	right *Node
}

/* Accessors and mutators for the Node class */
func (n *Node) Left() *Node {
	return n.left
}

func (n *Node) SetLeft(v *Node) {
	n.left = v
}

func (n *Node) Right() *Node {
	return n.right
}

func (n *Node) SetRight(v *Node) {
	n.right = v
}

func (n *Node) Value() int {
	return n.value
}

// Traversals

/*
	in order traversal
	L V R
*/
func InOrderTraversal(root *Node) {
	if root == nil {
		return
	}
	InOrderTraversal(root.Left())
	fmt.Println(root.Value())
	InOrderTraversal(root.Right())
}

/*
	pre order traversal
	V L R
*/
func PreOrderTraversal(root *Node) {
	if root == nil {
		return
	}
	fmt.Println(root.Value())
	PreOrderTraversal(root.Left())
	PreOrderTraversal(root.Right())
}

/*
	pre order traversal
	V R L
*/
func PostOrderTraversal(root *Node) {
	if root == nil {
		return
	}
	fmt.Println(root.Value())
	PostOrderTraversal(root.Right())
	PostOrderTraversal(root.Left())
}

// find the maximum depth of a tree
func MaxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	lDepth := MaxDepth(root.Left())
	rDepth := MaxDepth(root.Right())
	if lDepth > rDepth {
		return lDepth + 1
	}
	return rDepth + 1
}

// count number of nodes in a tree
func CountNodes(root *Node) int {
	if root == nil {
		return 0
	}
	return 1 + CountNodes(root.Left()) + CountNodes(root.Right())
}

func InsertValue(root *Node, value int) {
	if root == nil {
		root = &Node{value: value}
		return
	}

	if value <= root.Value() {
		if root.Left() == nil {
			root.SetLeft(&Node{value: value})
			return
		}
		InsertValue(root.Left(), value)
	} else {
		if root.Right() == nil {
			root.SetRight(&Node{value: value})
			return
		}
		InsertValue(root.Right(), value)
	}
}

/*
	data struct ascii viz
				 +1+
			+3+		 5
		 2	 4
*/
func CreateTestData() *Node {
	root := Node{value: 5}
	i := 0
	for i <= 10 {
		InsertValue(&root, i)
		i = i + 1
	}
	return &root
}

/*
func main() {
	root := createTestData()
	fmt.Println("pre")
	preOrderTraversal(root)
	fmt.Println("in")
	inOrderTraversal(root)
	fmt.Println("post")
	postOrderTraversal(root)
	fmt.Println("node count: ", countNodes(root))
	fmt.Println("max depth: ", maxDepth(root))
	// now insert a new value
	insertValue(root, 7)
	inOrderTraversal(root)
	fmt.Println("node count: ", countNodes(root))
	fmt.Println("max depth: ", maxDepth(root))
}*/
