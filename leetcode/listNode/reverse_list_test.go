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
	reverseLn := reverseList(&ln)
	fmt.Println(reverseLn.Val)

	reverseLn2 := reverseList2(&ln)
	fmt.Println(reverseLn2.Val)

}
