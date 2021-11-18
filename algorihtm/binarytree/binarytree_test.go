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
	isB, _ := isBalance(root)
	return isB
}

//返回当前树是否平衡二叉树，和当前树的深度
func isBalance(root *TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	}

	left, leftDepth := isBalance(root.Left)
	right, rightDepth := isBalance(root.Right)

	depth := leftDepth
	if leftDepth < rightDepth {
		depth = rightDepth
	}

	return left && right && math.Abs(float64(leftDepth-rightDepth)) <= 1, depth + 1
}

func TestIsBalance(t *testing.T) {
	root := NewTree2()
	isB := isBalanceTree(root)
	fmt.Println(isB)
}

/*
对称二叉树

给定一个二叉树，检查它是否是镜像对称的。

例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

*/

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	leftTree := root.Left
	rightTree := root.Right

	isSym := symmetricNode(leftTree, rightTree)
	return isSym
}

func symmetricNode(leftTree *TreeNode, rightTree *TreeNode) bool {
	if leftTree == nil && rightTree == nil {
		return true
	}
	if leftTree == nil || rightTree == nil {
		return false
	}
	return leftTree.Val == rightTree.Val && symmetricNode(leftTree.Left, rightTree.Right) && symmetricNode(leftTree.Right, rightTree.Left)
}

/*
给定一个二叉树，找出其最小深度。

最小深度是从根节点到最近叶子节点的最短路径上的节点数量。

说明：叶子节点是指没有子节点的节点。

*/

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}

	lD, rD := math.MaxInt32, math.MaxInt32
	if root.Left != nil {
		lD = minDepth(root.Left)
	}

	if root.Right != nil {
		rD = minDepth(root.Right)
	}

	if lD > rD {
		return rD + 1
	}
	return lD + 1
}

/*
翻转一棵二叉树。
输入：

     4
   /   \
  2     7
 / \   / \
1   3 6   9
输出：

     4
   /   \
  7     2
 / \   / \
9   6 3   1

*/
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	temp := root.Left
	root.Left = root.Right
	root.Right = temp

	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

func TestInvertTree(t *testing.T) {
	root := &TreeNode{Val: 4}
	root.Left = CreatTreeNode(2)
	root.Right = CreatTreeNode(7)

	root.Left.Left = CreatTreeNode(1)
	root.Left.Right = CreatTreeNode(3)

	root.Right.Left = CreatTreeNode(6)
	root.Right.Right = CreatTreeNode(9)

	root = invertTree(root)
	fmt.Println(root)
}

func CreatTreeNode(val int) *TreeNode {
	return &TreeNode{
		Val: val,
	}
}

/*
给定两个二叉树，想象当你将它们中的一个覆盖到另一个上时，两个二叉树的一些节点便会重叠。

你需要将他们合并为一个新的二叉树。合并的规则是如果两个节点重叠，那么将他们的值相加作为节点合并后的新值，否则不为NULL 的节点将直接作为新二叉树的节点。

示例1:

输入:
	Tree 1                     Tree 2
          1                         2
         / \                       / \
        3   2                     1   3
       /                           \   \
      5                             4   7
输出:
合并后的树:
	     3
	    / \
	   4   5
	  / \   \
	 5   4   7

注意:合并必须从两个树的根节点开始。

*/

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	root := &TreeNode{}
	if root1 == nil {
		return root2
	} else if root2 == nil {
		return root1
	} else {
		root.Val = root1.Val + root2.Val
		root.Left = mergeTrees(root1.Left, root2.Left)
		root.Right = mergeTrees(root1.Right, root2.Right)
	}
	return root
}
