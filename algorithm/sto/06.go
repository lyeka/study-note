package sto


type ListNode struct {
      Val int
      Next *ListNode
}

func reversePrint(head *ListNode) []int {
	out := []int{}
	if head == nil {
		return out
	}
	reverse(head, &out)
	return out
}

func reverse(head *ListNode, out *[]int){
	if head == nil {
		return
	}
	reverse(head.Next, out)
	*out = append(*out, head.Val)
}
