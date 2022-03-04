package __jianzhioffer

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Value int
}

//前序（中左右） 中序（左中右） 构造二叉树
//前序遍历序列{1,2,4,7,3,5,6,8}
//中序遍历序列{4,7,2,1,5,3,8,6}
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	head := &TreeNode{Value: preorder[0]}
	mid := findMid(head.Value, inorder)
	head.Left = buildTree(preorder[1:mid+1], inorder[:mid])
	head.Right = buildTree(preorder[mid+1:], inorder[mid+1:])

	return head
}

func findMid(val int, in []int) (index int) {
	for i, v := range in {
		if v == val {
			return i
		}
	}
	return
}
