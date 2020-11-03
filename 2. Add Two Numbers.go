package main

import (
	shared "github.com/santoshpanna/leetcode/shared"
)

func add(val int, node *shared.ListNode) *shared.ListNode {
	var temp = &shared.ListNode{
		Val: val,
		Next: nil,
	}

	if node == nil {
		node = temp
	} else {
		var itr = node
		for itr.Next != nil {
			itr = itr.Next
		}
		itr.Next = temp
	}

	return node
}


func addTwoNumbers(l1 *shared.ListNode, l2 *shared.ListNode) *shared.ListNode {
	var res *shared.ListNode
	var carry int = 0

	for l1 != nil || l2 != nil {
		if l1 != nil {
			carry = carry + l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			carry = carry + l2.Val
			l2 = l2.Next
		}

		res = add(carry % 10, res)
		carry = carry / 10
	}
	if carry > 0 {
		res = add(carry % 10, res)
	}

	return res
}

func main() {
	var l1 *shared.ListNode
	var l2 *shared.ListNode

	l1 = shared.MakeSLL([] int {5,6,4})
	l2 = shared.MakeSLL([] int {2,4,3})

	var result = addTwoNumbers(l1, l2)

	shared.PrintSSL(result)

}
