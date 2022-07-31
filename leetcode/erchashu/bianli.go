package erchashu

//链接：https://leetcode-cn.com/problems/binary-tree-inorder-traversal/solution/er-cha-shu-de-zhong-xu-bian-li-by-leetcode-solutio/

// 前序遍历：先访问根节点，再前序遍历左子树，再前序遍历右子树
// 中序遍历：先中序遍历左子树，再访问根节点，再中序遍历右子树
// 后序遍历：先后序遍历左子树，再后序遍历右子树，再访问根节点

func preorderTraversal(root *TreeNode) (res []int) {
	if root == nil {
		return
	}
	// 先访问根再访问左右
	res = append(res, root.Val)
	preorderTraversal(root.Left)
	preorderTraversal(root.Right)
	return
}
func preorderTraversal2(root *TreeNode) (res []int) {
	// 非递归
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)

	for root != nil || len(stack) != 0 {
		for root != nil {
			// 前序遍历，所以先保存结果
			result = append(result, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		// pop
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = node.Right
	}
	return result
}

func inorderTraversal(root *TreeNode) (res []int) {
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		res = append(res, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return
}
