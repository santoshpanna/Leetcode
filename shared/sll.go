package shared

import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func add (val int, listNode *ListNode) *ListNode {
	var temp = &ListNode{
		Val: val,
		Next: nil,
	}

	if listNode == nil {
		listNode = temp
	} else {
		var itr = listNode
		for itr.Next != nil {
			itr = itr.Next
		}
		itr.Next = temp
	}

	return listNode
}

func MakeSLL(arr[] int) *ListNode {
	var listNode *ListNode
	for _, item := range arr {
		listNode = add(item, listNode)
	}

	return listNode
}

func Length(node *ListNode) int {
	var length int = 0

	for ; node != nil; node = node.Next {
		length += 1
	}

	return length
}

func PrintSSL(node *ListNode) {
	fmt.Print("\nLinked List : ")

	for ; node != nil; node = node.Next {
		fmt.Print(node.Val, " -> ")
	}
	fmt.Print("\n")

}

func main()  {
	fmt.Println("make with array == ", MakeSLL([]int{2,3,4}))
}