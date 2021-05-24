package binarytree

/*
给定一个二叉树，找出其最大深度。

二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//方法一：递归 深度优先遍历 dfs
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}

//方法二：广度优先遍历
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	nodeList := make([]*TreeNode, 0)
	nodeList = append(nodeList, root)
	depth := 0
	for len(nodeList) > 0 {

		sz := len(nodeList)
		for sz > 0 {
			cruNode := nodeList[0]
			nodeList = nodeList[1:]
			if cruNode.Left != nil {
				nodeList = append(nodeList, cruNode.Left)
			}
			if cruNode.Right != nil {
				nodeList = append(nodeList, cruNode.Right)
			}
			sz--
		}
		depth++
	}
	return depth
}
