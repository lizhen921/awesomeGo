package binarytree

import (
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

//左叶子之和
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return 0
	}
	sum := 0
	stack := make([]*TreeNode, 0)
	isLeftNode := make([]bool, 0)
	for root != nil {
		stack = append(stack, root)
		isLeftNode = append(isLeftNode, true)
		root = root.Left
	}

	for len(stack) != 0 {
		root = stack[len(stack)-1]
		isLeft := isLeftNode[len(stack)-1]

		stack = stack[:len(stack)-1]
		isLeftNode = isLeftNode[:len(isLeftNode)-1]

		if isLeft && root.Left == nil && root.Right == nil {
			sum += root.Val
		}

		root = root.Right
		if root != nil {
			stack = append(stack, root)
			isLeftNode = append(isLeftNode, false)
			root = root.Left
		}
		for root != nil {
			stack = append(stack, root)
			isLeftNode = append(isLeftNode, true)
			root = root.Left
		}
	}
	return sum
}

//路径之和
func hasPathSum001(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false

	}
	if root.Left == nil && root.Right == nil && targetSum == 0 {
		return true
	}

	isLeft := hasPathSum001(root.Left, targetSum-root.Val)
	isRight := hasPathSum001(root.Right, targetSum-root.Val)

	return isLeft || isRight
}

func hasPathSum002(root *TreeNode, targetSum int) bool {
	stack := make([]*TreeNode, 0)
	tempSumStack := make([]int, 0)

	tempFather := 0
	for root != nil {
		stack = append(stack, root)
		tempSumStack = append(tempSumStack, tempFather+root.Val)
		if root.Left == nil && root.Right == nil && tempFather+root.Val == targetSum {
			return true
		}
		tempFather = tempFather + root.Val
		root = root.Left
	}

	for len(stack) != 0 {
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		tempFather = tempSumStack[len(tempSumStack)-1]
		tempSumStack = tempSumStack[:len(tempSumStack)-1]

		root = root.Right
		for root != nil {
			stack = append(stack, root)
			tempSumStack = append(tempSumStack, tempFather+root.Val)
			if root.Left == nil && root.Right == nil && tempFather+root.Val == targetSum {
				return true
			}
			root = root.Left
		}
	}

	return false
}

/*
给定一个不含重复元素的整数数组 nums 。一个以此数组直接递归构建的 最大二叉树 定义如下：

二叉树的根是数组 nums 中的最大元素。
左子树是通过数组中 最大值左边部分 递归构造出的最大二叉树。
右子树是通过数组中 最大值右边部分 递归构造出的最大二叉树。
返回有给定数组 nums 构建的 最大二叉树 。

*/

func constructMaximumBinaryTree001(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	index := getMax(nums)
	root := &TreeNode{Val: nums[index]}
	root.Left = constructMaximumBinaryTree001(nums[:index])
	root.Right = constructMaximumBinaryTree001(nums[index+1:])

	return root
}

func getMax(nums []int) int {
	index := 0
	max := nums[0]

	for i, v := range nums {
		if max < v {
			index = i
			max = v
		}
	}
	return index
}

/*
preorder = [3,9,20,15,7]  中左右 [1,2]
inorder = [9,3,15,20,7]   左中右 [1,2]

Output: [3,9,20,null,null,15,7]

*/
func buildTree001(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	index := getIndex(preorder[0], inorder)

	root := &TreeNode{Val: preorder[0]}
	root.Left = buildTree001(preorder[1:index+1], inorder[:index])
	root.Right = buildTree001(preorder[index+1:], inorder[index+1:])

	return root
}

func getIndex(value int, inorder []int) int {
	for i, v := range inorder {
		if v == value {
			return i
		}
	}
	return 0
}

/*
inorder = [9,3,15,20,7] 左中右
postorder = [9,15,7,20,3] 左右中

Output: [3,9,20,null,null,15,7]
*/
func buildTree002(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}

	index := getIndex(postorder[len(postorder)-1], inorder)

	root := &TreeNode{Val: postorder[len(postorder)-1]}
	root.Left = buildTree002(inorder[:index], postorder[:index])
	root.Right = buildTree002(inorder[index+1:], postorder[index:len(postorder)-1])

	return root
}

func mergeTrees001(root1 *TreeNode, root2 *TreeNode) *TreeNode {

	root := &TreeNode{}

	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	root.Val = root1.Val + root2.Val
	root.Left = mergeTrees001(root1.Left, root2.Left)
	root.Right = mergeTrees001(root1.Right, root2.Right)

	return root
}

func searchBST001(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	if root.Val > val {
		root = root.Left
	} else {
		root = root.Right
	}
	return searchBST001(root, val)
}

func searchBST002(root *TreeNode, val int) *TreeNode {

	for root != nil && root.Val != val {
		if root.Val > val {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return root

}

func TestInBinaryTree(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Left = CreatTreeNode(2)
	root.Left.Left = CreatTreeNode(3)
	root.Left.Right = CreatTreeNode(4)
	root.Left.Right.Left = CreatTreeNode(5)
	root.Left.Right.Right = CreatTreeNode(6)

	root.Right = CreatTreeNode(2)
	root.Right.Right = CreatTreeNode(3)
	root.Right.Left = CreatTreeNode(4)

	//isSymmetric001(root)
	//res := preBinaryTree(root)
	//fmt.Println(res)
	//res = inBinaryTree(root)
	//fmt.Println(res)
	//res = postBinaryTree(root)
	//fmt.Println(res)
	//fmt.Println(binaryTreePaths002(root))

	//hasPathSum002(root, 13)

	//constructMaximumBinaryTree001([]int{3, 2, 1, 6, 0, 5})
	buildTree001([]int{1, 2}, []int{1, 2})
}
