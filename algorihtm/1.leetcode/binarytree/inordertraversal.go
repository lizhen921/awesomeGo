package binarytree

/*
给定一个二叉树的根节点 root ，返回它的 中序 遍历。
实例1：
输入：root = [1,null,2,3]
输出：[1,3,2]
示例 2：

输入：root = []
输出：[]
示例 3：

输入：root = [1]
输出：[1]
*/

var preOrder = make([]int, 0)
var inOrder = make([]int, 0)
var afterOrder = make([]int, 0)

// 递归法
func inorderTraversal(root *TreeNode) ([]int, []int, []int) {
	if root == nil {
		return preOrder, inOrder, afterOrder
	}
	preOrder = append(preOrder, root.Val)
	inorderTraversal(root.Left)
	inOrder = append(inOrder, root.Val)
	inorderTraversal(root.Right)
	afterOrder = append(afterOrder, root.Val)
	return preOrder, inOrder, afterOrder
}

//迭代法 -- 中序遍历 左中右
func inorderTraversal2(root *TreeNode) []int {

	stack := make([]*TreeNode, 0)
	res := make([]int, 0)

	for root != nil {
		stack = append(stack, root)
		root = root.Left
	}

	for len(stack) != 0 {
		stackLen := len(stack)
		needPop := stack[stackLen-1]   //获取栈顶元素
		res = append(res, needPop.Val) //记录出栈元素
		stack = stack[:stackLen-1]     //出栈
		root = needPop.Right
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
	}
	return res
}

//迭代法 -- 前序遍历 中左右
func inorderTraversal3(root *TreeNode) []int {

	stack := make([]*TreeNode, 0)
	res := make([]int, 0)

	for root != nil {
		res = append(res, root.Val)
		stack = append(stack, root)
		root = root.Left
	}

	for len(stack) != 0 {
		stackLen := len(stack)
		needPop := stack[stackLen-1] //获取栈顶元素
		stack = stack[:stackLen-1]   //出栈

		root = needPop.Right
		for root != nil {
			res = append(res, root.Val) //记录入栈元素
			stack = append(stack, root)
			root = root.Left

		}
	}
	return res
}

//迭代法 -- 后续遍历 左右中
func inorderTraversal4(root *TreeNode) []int {

	stack := make([]*TreeNode, 0)
	res := make([]int, 0)

	pre := &TreeNode{}
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		stackLen := len(stack)
		root = stack[stackLen-1]   //获取栈顶元素
		stack = stack[:stackLen-1] //出栈
		if root.Right == nil || pre == root.Right {
			res = append(res, root.Val)
			pre = root
			root = nil
		} else {
			stack = append(stack, root)
			root = root.Right
		}
	}
	return res
}
