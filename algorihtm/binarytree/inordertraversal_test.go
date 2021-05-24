package binarytree

import (
	"fmt"
	"testing"
)

//递归
func TestInOrderTrancersal(t *testing.T) {
	root := NewTree()
	pre, in, after := inorderTraversal(root)
	fmt.Println("前序：", pre)
	fmt.Println("中序：", in)
	fmt.Println("后序：", after)
}

//递归-中序
func TestInOrderTrancersal2(t *testing.T) {
	root := NewTree()
	res := inorderTraversal2(root)
	fmt.Println(res)
}

//递归-前序
func TestInOrderTrancersa3(t *testing.T) {
	root := NewTree()
	res := inorderTraversal3(root)
	fmt.Println(res)
}

//递归-后序
func TestInOrderTrancersa4(t *testing.T) {
	root := NewTree()
	res := inorderTraversal4(root)
	fmt.Println(res)
}
