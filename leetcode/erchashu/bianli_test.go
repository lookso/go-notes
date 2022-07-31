package erchashu

import "testing"

var tn = TreeNode{
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

func TestErChaBianLi(t *testing.T) {
	t.Log("前序遍历:", preorderTraversal(&tn))
	t.Log("中序遍历:", inorderTraversal(&tn))
}

func BenchmarkErChaBianLi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.Log("前序遍历:", preorderTraversal(&tn))
	}
}
