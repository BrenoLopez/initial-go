package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	var newList *ListNode
	tail := &newList
	for l1 != nil || l2 != nil || carry != 0 {
		x, y := 0, 0

		if l1 != nil {
			x = l1.Val
		}
		if l2 != nil {
			y = l2.Val
		}

		sum := x + y + carry
		carry = sum / 10

		*tail = &ListNode{Val: sum % 10} // altera o valor da desreferencia de tail o que altera o valor de newList
		tail = &(*tail).Next             // atualiza tail com o valor do proximo espaço de memoria
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	return newList
}

func main() {
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}

	addTwoNumbers(l1, l2)

}
