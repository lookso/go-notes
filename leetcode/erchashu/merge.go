package erchashu

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	// 遇到一颗树为空了,那不用递归了，返回非空的就完事了
	if t2 == nil {
		return t1
	}
	if t1 == nil {
		return t2
	}
	// 相加取值
	t1.Val += t2.Val
	// 左右递归
	t1.Left = mergeTrees(t1.Left, t2.Left)
	t1.Right = mergeTrees(t1.Right, t2.Right)
	return t1
}

//作者：feng-a
//链接：https://leetcode-cn.com/problems/merge-two-binary-trees/solution/617-he-bing-er-cha-shu-di-gui-jian-ji-ming-liao-by/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
