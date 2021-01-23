package erchashu

import (
	"testing"
)

func TestIsValidBST(t *testing.T) {
	tn := TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 10,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	t.Log("是否是有效二叉搜索树:", isValidBST(&tn))
}
