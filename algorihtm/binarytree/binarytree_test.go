package binarytree

/*
给定一个二叉树，找出其最大深度。

二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

说明: 叶子节点是指没有子节点的节点。

示例：
给定二叉树 [3,9,20,null,null,15,7]，

*/
func maxDepth0(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDeep := maxDepth0(root.Left)
	rightDeep := maxDepth0(root.Right)
	if leftDeep > rightDeep {
		return leftDeep + 1
	}

	return rightDeep + 1
}
