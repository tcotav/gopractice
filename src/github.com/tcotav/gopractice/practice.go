package main

import (
	"fmt"
	"github.com/tcotav/gopractice/bst"
)

func main() {

	root := bst.CreateTestData()
	fmt.Println("pre")
	bst.PreOrderTraversal(root)
	fmt.Println("in")
	bst.InOrderTraversal(root)
	fmt.Println("post")
	bst.PostOrderTraversal(root)
	fmt.Println("node count: ", bst.CountNodes(root))
	fmt.Println("max depth: ", bst.MaxDepth(root))
	// now insert a new value
	bst.InsertValue(root, 7)
	bst.InOrderTraversal(root)
	fmt.Println("node count: ", bst.CountNodes(root))
	fmt.Println("max depth: ", bst.MaxDepth(root))

}
