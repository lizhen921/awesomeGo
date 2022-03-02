package binarytree

import (
	"fmt"
	"testing"
)

func TestMaxDepth(t *testing.T) {
	root := NewTree()
	depth := maxDepth(root)
	fmt.Println(depth)
}

func TestMaxDepth2(t *testing.T) {
	root := NewTree()
	depth := maxDepth2(root)
	fmt.Println(depth)
}

func NewTree() *TreeNode {
	n1 := &TreeNode{1, nil, nil}
	n2 := &TreeNode{2, nil, nil}
	n3 := &TreeNode{3, nil, nil}
	n4 := &TreeNode{4, nil, nil}
	n5 := &TreeNode{5, nil, nil}
	n6 := &TreeNode{6, nil, nil}
	n7 := &TreeNode{7, nil, nil}
	n8 := &TreeNode{8, nil, nil}
	n9 := &TreeNode{9, nil, nil}
	n10 := &TreeNode{10, nil, nil}
	n11 := &TreeNode{11, nil, nil}
	n1.Left = n2
	n1.Right = n3
	n2.Left = n4
	n3.Left = n5
	n3.Right = n6
	n5.Left = n7
	n6.Right = n8
	n7.Left = n9
	n7.Right = n10
	n10.Left = n11

	return n1
}
func NewTree2() *TreeNode {
	n1 := &TreeNode{1, nil, nil}
	n2 := &TreeNode{2, nil, nil}
	n3 := &TreeNode{3, nil, nil}
	n4 := &TreeNode{4, nil, nil}
	n5 := &TreeNode{5, nil, nil}
	n6 := &TreeNode{6, nil, nil}
	n7 := &TreeNode{7, nil, nil}
	n8 := &TreeNode{8, nil, nil}
	n9 := &TreeNode{9, nil, nil}
	n10 := &TreeNode{10, nil, nil}
	n11 := &TreeNode{11, nil, nil}
	n1.Left = n2
	n1.Right = n3
	n2.Left = n4
	n2.Right = n5
	n3.Left = n6
	n3.Right = n7
	n4.Left = n8
	n4.Right = n9
	n6.Left = n10
	n6.Right = n11

	return n1
}
