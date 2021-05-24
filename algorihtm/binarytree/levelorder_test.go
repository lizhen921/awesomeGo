package binarytree

import (
	"fmt"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	tree := NewTree()
	listTree := levelOrder(tree)
	fmt.Println(listTree)
}
