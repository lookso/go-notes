package listNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var temp *ListNode
	for head != nil {
		temp = &ListNode{Val: head.Val, Next: temp}
		head = head.Next
	}
	return temp
}
func reverseList2(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode = nil
	for cur != nil {
		pre, cur, cur.Next = cur, cur.Next, pre
	}
	return pre
}

//链接：https://leetcode-cn.com/problems/fan-zhuan-lian-biao-lcof/solution/wo-ye-bu-zhi-dao-zen-yao-jiu-shuang-bai-liao-by-ea/
