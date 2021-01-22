package erchashu

//https://leetcode-cn.com/problems/binary-tree-level-order-traversal/solution/

//二叉树：[3,9,20,null,null,15,7],
//
//3
/// \
//9  20
//  /  \
// 15   7
//返回其层序遍历结果：
//
//[[3],
//[9,20],
//[15,7]]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/binary-tree-level-order-traversal
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func levelOrder(root *TreeNode) [][]int {
	return dfs(root, 0, [][]int{})
}

func dfs(root *TreeNode, level int, res [][]int) [][]int {
	if root == nil {
		return res
	}
	if len(res) == level {
		res = append(res, []int{root.Val})
	} else {
		res[level] = append(res[level], root.Val)
	}
	res = dfs(root.Left, level+1, res)
	res = dfs(root.Right, level+1, res)
	return res
}

// [3,9,20,null,null,15,7]

