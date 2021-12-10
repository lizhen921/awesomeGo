package binarytree

/*

 */
/*

 */

func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return nil
	}
	res = append(res, root.Val)
	res = append(res, preorderTraversal(root.Left)...)
	res = append(res, preorderTraversal(root.Right)...)
	return res
}
