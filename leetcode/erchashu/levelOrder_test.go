package erchashu

import (
	"testing"
)

func TestLevelOrder(t *testing.T) {
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
	t.Log("层级遍历:",levelOrder(&tn))
}
