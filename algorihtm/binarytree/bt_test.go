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

//迭代 对称二叉树
func isSymmetric001(root *TreeNode) bool {
	leftStack := make([]*TreeNode, 0)
	rightStack := make([]*TreeNode, 0)
	left, right := root, root

	for left != nil || right != nil {
		if left == nil || right == nil {
			return false
		}
		leftStack = append(leftStack, left)
		rightStack = append(rightStack, right)
		if left.Val != right.Val {
			return false
		}
		left = left.Left
		right = right.Right
	}

	for len(leftStack) != 0 && len(rightStack) != 0 {
		left = leftStack[len(leftStack)-1]
		right = rightStack[len(rightStack)-1]
		leftStack = leftStack[:len(leftStack)-1]
		rightStack = rightStack[:len(rightStack)-1]
		left = left.Right
		right = right.Left
		for left != nil || right != nil {
			if left == nil || right == nil {
				return false
			}
			leftStack = append(leftStack, left)
			rightStack = append(rightStack, right)
			if left.Val != right.Val {
				return false
			}
			left = left.Left
			right = right.Right
		}
	}

	return true
}

//递归 对称二叉树
func isSymmetric002(root *TreeNode) bool {
	if root == nil {
		return true
	}
	left, right := root.Left, root.Right
	return SameValue(left, right)
}

func SameValue(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	leftEqual := SameValue(left.Left, right.Right)
	rightEqual := SameValue(left.Right, right.Left)
	return leftEqual && rightEqual && (left.Val == right.Val)
}

//最大深度，最小深度
func MaxDepth001(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depth := 0
	nodeList := make([]*TreeNode, 0)
	nodeList = append(nodeList, root)
	levelCount := 1
	for len(nodeList) != 0 {
		tempCount := 0
		for i := 0; i < levelCount; i++ {
			node := nodeList[0]
			nodeList = nodeList[1:]

			if node.Left != nil {
				nodeList = append(nodeList, node.Left)
				tempCount++
			}

			if node.Right != nil {
				nodeList = append(nodeList, node.Right)
				tempCount++
			}

		}
		levelCount = tempCount
		depth++
	}

	return depth
}

func MinDepth001(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depth := 0
	nodeList := make([]*TreeNode, 0)
	nodeList = append(nodeList, root)
	levelCount := 1
	for len(nodeList) != 0 {
		depth++

		tempCount := 0
		for i := 0; i < levelCount; i++ {
			node := nodeList[0]
			nodeList = nodeList[1:]

			if node.Left == nil && node.Right == nil {
				return depth
			}

			if node.Left != nil {
				nodeList = append(nodeList, node.Left)
				tempCount++
			}

			if node.Right != nil {
				nodeList = append(nodeList, node.Right)
				tempCount++
			}

		}
		levelCount = tempCount

	}

	return depth
}
func TestInBinaryTree(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Left = CreatTreeNode(2)
	root.Left.Left = CreatTreeNode(3)
	root.Left.Right = CreatTreeNode(4)

	root.Right = CreatTreeNode(2)
	root.Right.Right = CreatTreeNode(3)
	root.Right.Left = CreatTreeNode(4)

	isSymmetric001(root)
	//res := preBinaryTree(root)
	//fmt.Println(res)
	//res = inBinaryTree(root)
	//fmt.Println(res)
	//res = postBinaryTree(root)
	//fmt.Println(res)
}
