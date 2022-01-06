package binarytree

import (
	"fmt"
	"math"
	"strconv"
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

//最大深度，
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

func MaxDepth002(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := MaxDepth002(root.Left)
	rightDepth := MaxDepth002(root.Right)
	return max(leftDepth, rightDepth) + 1
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

//最小深度
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

func MinDepth002(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 0
	}
	leftDepth := MinDepth002(root.Left)
	rightDepth := MinDepth002(root.Right)

	if root.Left == nil || root.Right == nil {

		return max(leftDepth, rightDepth) + 1
	}
	return min(leftDepth, rightDepth) + 1
}

func min(a, b int) int {
	if a >= b {
		return b
	}
	return a
}

//左右差不超过1
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	isB, _ := isBalancedChild(root)

	return isB

}

func isBalancedChild(root *TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	}
	isLeftBalanced, depthLeft := isBalancedChild(root.Left)
	isRightBalanced, depthRight := isBalancedChild(root.Right)

	depth := depthLeft
	if depthRight > depthLeft {
		depth = depthRight
	}
	return isLeftBalanced && isRightBalanced && math.Abs(float64(depthLeft-depthRight)) <= 1, depth + 1
}

//路径
func binaryTreePaths001(root *TreeNode) []string {
	paths := make([]string, 0)

	queue := make([]*TreeNode, 0)
	queuePath := make([]string, 0)

	queue = append(queue, root)
	queuePath = append(queuePath, strconv.Itoa(root.Val))

	num := 1
	for len(queue) != 0 {
		tempNum := num
		num = 0

		for i := 0; i < tempNum; i++ {
			root := queue[0]
			rootPath := queuePath[0]

			queue = queue[1:]
			queuePath = queuePath[1:]

			if root.Left != nil {
				queue = append(queue, root.Left)
				queuePath = append(queuePath, rootPath+"->"+strconv.Itoa(root.Left.Val))
				num++
			}

			if root.Right != nil {
				queue = append(queue, root.Right)
				queuePath = append(queuePath, rootPath+"->"+strconv.Itoa(root.Right.Val))
				num++
			}
			if root.Left == nil && root.Right == nil {
				paths = append(paths, rootPath)
			}
		}
	}
	return paths
}

//路径
func binaryTreePaths002(root *TreeNode) []string {
	res := make([]string, 0)
	if root == nil {
		return res
	}
	buildPath(root.Left, strconv.Itoa(root.Val), &res)

	buildPath(root.Right, strconv.Itoa(root.Val), &res)

	return res
}

func buildPath(root *TreeNode, pathStr string, paths *[]string) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		pathStr = pathStr + "->" + strconv.Itoa(root.Val)
		*paths = append(*paths, pathStr)
		return
	}

	pathStr = pathStr + "->" + strconv.Itoa(root.Val)

	buildPath(root.Left, pathStr, paths)
	buildPath(root.Right, pathStr, paths)
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
	fmt.Println(binaryTreePaths002(root))
}
