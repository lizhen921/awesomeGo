package binarytree

import (
	"fmt"
	"math"
	"testing"
)

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

/*
给定一个二叉树，判断它是否是高度平衡的二叉树。

本题中，一棵高度平衡二叉树定义为：

一个二叉树每个节点的左右两个子树的高度差的绝对值不超过 1 。
*/

//自上而下 递归
/*
结束条件： root==nil
返回值：left right  depth<1
*/
func isBalanceTree(root *TreeNode) bool {
	isB,_ := isBalance(root)
	return isB
}
//返回当前树是否平衡二叉树，和当前树的深度
func isBalance(root *TreeNode) (bool,int) {
	if root == nil{
		return true,0
	}

	left, leftDepth := isBalance(root.Left)
	right,rightDepth := isBalance(root.Right)

	depth := leftDepth
	if leftDepth < rightDepth {
		depth = rightDepth
	}

	return left && right && math.Abs(float64(leftDepth-rightDepth))<=1 , depth+1
}


//自下而上 递归
func isBalanceTree2(root *TreeNode) bool {

	return false
}

func TestIsBalance(t *testing.T)  {
	root := NewTree2()
	isB := isBalanceTree(root)
	fmt.Println(isB)
}