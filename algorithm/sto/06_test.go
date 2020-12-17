package sto

import (
	"fmt"
	"testing"
)

func Test06(t *testing.T) {
	ln := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 2,
				Next: nil,
			},
		},
	}

	tmp := []int{2, 3, 1}
	out := reversePrint(ln)
	if len(tmp) != len(out) {
		t.Fail()
	}
	for i, v := range out{
		if tmp[i] != v {
			fmt.Println("out: ",out)
			t.Fail()
		}
	}
}
