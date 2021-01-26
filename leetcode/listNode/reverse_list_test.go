package listNode

import (
	"fmt"
	"testing"
)

func TestReverseList(t *testing.T) {
	var ln = ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	reverLn:=reverseList(&ln)
	fmt.Println(reverLn.Val)
}
