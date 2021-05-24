package binarytree

/*
给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。
示例：
二叉树：[3,9,20,null,null,15,7],
3
/ \
9  20
/  \
15   7
返回其层序遍历结果：
[
[3],
[9,20],
[15,7]
]
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-level-order-traversal
*/

func levelOrder(root *TreeNode) [][]int {
	var listTree = make([][]int, 0)
	if root == nil {
		return listTree
	}
	tempList := make([]*TreeNode, 0)
	tempList = append(tempList, root)
	for len(tempList) > 0 {
		sz := len(tempList)
		listTreeLine := make([]int, 0)
		for sz > 0 {
			curNode := tempList[0]
			listTreeLine = append(listTreeLine, curNode.Val)
			tempList = tempList[1:]
			if curNode.Left != nil {
				tempList = append(tempList, curNode.Left)
			}
			if curNode.Right != nil {
				tempList = append(tempList, curNode.Right)
			}
			sz--
		}
		listTree = append(listTree, listTreeLine)
	}
	return listTree
}
