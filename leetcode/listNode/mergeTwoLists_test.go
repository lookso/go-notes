package listNode

import (
	"fmt"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	var a = ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	var b = ListNode{
		Val: 6,
		Next: &ListNode{
			Val: 7,
			Next: &ListNode{
				Val:  9,
				Next: nil,
			},
		},
	}
	p := mergeTwoLists(&a, &b)
	for p != nil {
		fmt.Println(*p)
		p = p.Next //移动指针
	}
}
