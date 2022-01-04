package add_two_numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	nowL1 := l1.Next
	nowL2 := l2.Next
	headNode := ListNode{Val: l1.Val + l2.Val, Next: nil}
	carry := 0
	if headNode.Val > 9 {
		carry = 1
		headNode.Val -= 10
	}
	tailNode := &headNode
	val := 0
	for nowL1 != nil || nowL2 != nil {
		if nowL1 == nil {
			val = nowL2.Val + carry
		} else if nowL2 == nil {
			val = nowL1.Val + carry
		} else {
			val = nowL1.Val + nowL2.Val + carry
		}
		if val > 9 {
			carry = 1
			val -= 10
		} else {
			carry = 0
		}
		newNode := ListNode{Val: val, Next: nil}
		tailNode.Next = &newNode
		tailNode = &newNode
		if nowL1 != nil {
			nowL1 = nowL1.Next
		}
		if nowL2 != nil {
			nowL2 = nowL2.Next
		}
	}
	if carry == 1 {
		newNode := ListNode{Val: 1, Next: nil}
		tailNode.Next = &newNode
	}
	return &headNode
}
