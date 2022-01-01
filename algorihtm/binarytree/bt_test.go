package binarytree

import (
	"testing"
)

/*

 */
func preOrderTree(root *TreeNode) []int {
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)

	for root != nil {
		stack = append(stack, root)
		res = append(res, root.Val)
		root = root.Left
	}

	for len(stack) != 0 {
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		root = root.Right
		for root != nil {
			stack = append(stack, root)
			res = append(res, root.Val)
			root = root.Left
		}
	}
	return res
}

/*
左中右
*/

func inOrderTree(root *TreeNode) []int {
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)

	for root != nil {
		stack = append(stack, root)
		root = root.Left
	}

	for len(stack) != 0 {
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)

		root = root.Right
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
	}
	return res
}

//后序 左右中
func postOrder(root *TreeNode) []int {
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)

	pre := &TreeNode{}
	for root != nil || len(stack) != 0 {

		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]

		if root.Right == nil || pre == root.Right {
			res = append(res, root.Val)
			stack = stack[:len(stack)-1]
			pre = root
			root = nil
		} else {
			root = root.Right
		}
	}

	return res
}

func levelOrderBtree(root *TreeNode) [][]int {
	res := make([][]int, 0)
	queue := make(chan *TreeNode, 9999)
	if root == nil {
		return res
	}
	queue <- root
	//level := 1
	count := 1
	for len(queue) != 0 {
		resLevel := make([]int, 0)
		for i := 0; i < count; i++ {
			root := <-queue
			resLevel = append(resLevel, root.Val)
			if root.Left != nil {
				queue <- root.Left
				count++
			}

			if root.Right != nil {
				queue <- root.Right
				count++
			}
		}
		res = append(res, resLevel)
	}
	return res
}

func reverseTree(root *TreeNode) *TreeNode {
	nodeList := make([]*TreeNode, 0)
	node := root
	nodeList = append(nodeList, node)

	num := 1
	for len(nodeList) != 0 {
		for i := 0; i < num; i++ {
			node = nodeList[0]
			nodeList = nodeList[1:]
			if node == nil {
				continue
			}
			if node.Left != nil || node.Right != nil {
				node.Left, node.Right = node.Right, node.Left
				nodeList = append(nodeList, node.Left)
				nodeList = append(nodeList, node.Right)
				num += 2
			}
		}
	}
	return root
}

func reverseTree2(root *TreeNode) *TreeNode {
	stack := make([]*TreeNode, 0)
	node := root

	for node != nil {
		stack = append(stack, node)
		node = node.Left
	}

	for len(stack) != 0 {
		node = stack[len(stack)-1]
		node.Left, node.Right = node.Right, node.Left
		stack = stack[:len(stack)-1]

		node = node.Left
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
	}
	return root
}

func TestInBinaryTree(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Right = CreatTreeNode(2)
	root.Right.Left = CreatTreeNode(3)

	//res := preBinaryTree(root)
	//fmt.Println(res)
	//res = inBinaryTree(root)
	//fmt.Println(res)
	//res = postBinaryTree(root)
	//fmt.Println(res)
}
