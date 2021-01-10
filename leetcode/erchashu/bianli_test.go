package erchashu

import "testing"

func TestErChaBianLi(t *testing.T) {
	tn := TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 25,
			},
		},
	}
	t.Log("前序遍历:", preorderTraversal(&tn))
	t.Log("中序遍历:", inorderTraversal(&tn))
}
