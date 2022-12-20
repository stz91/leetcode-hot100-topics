package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addNode(l **ListNode, val int) {
	if *l == nil {
		*l = &ListNode{
			Val: val,
		}
		return
	}
	curNode := *l
	for curNode.Next != nil {
		curNode = curNode.Next
	}
	curNode.Next = &ListNode{
		Val: val,
	}
}

func Traverse(l *ListNode) {
	if l == nil {
		return
	}
	curNode := l
	for curNode != nil {
		fmt.Printf("curNode value's %d\n", curNode.Val)
		curNode = curNode.Next
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	curNode1 := l1
	curNode2 := l2
	var resList *ListNode
	var lastQuotient int
	var tail *ListNode
	for curNode1 != nil && curNode2 != nil {
		curNode := &ListNode{}
		curNode.Val = (curNode1.Val + curNode2.Val + lastQuotient) % 10
		lastQuotient = (curNode1.Val + curNode2.Val + lastQuotient) / 10
		curNode1 = curNode1.Next
		curNode2 = curNode2.Next
		if tail == nil {
			tail = curNode
		} else {
			tail.Next = curNode
			tail = tail.Next
		}
		if resList == nil {
			resList = curNode
		}
	}
	for curNode1 != nil {
		curNode := &ListNode{}
		curNode.Val = (curNode1.Val + lastQuotient) % 10
		lastQuotient = (curNode1.Val + lastQuotient) / 10
		curNode1 = curNode1.Next
		if tail == nil {
			tail = curNode
		} else {
			tail.Next = curNode
			tail = tail.Next
		}
		if resList == nil {
			resList = curNode
		}
	}
	for curNode2 != nil {
		curNode := &ListNode{}
		curNode.Val = (curNode2.Val + lastQuotient) % 10
		lastQuotient = (curNode2.Val + lastQuotient) / 10
		curNode2 = curNode2.Next
		if tail == nil {
			tail = curNode
		} else {
			tail.Next = curNode
			tail = tail.Next
		}
		if resList == nil {
			resList = curNode
		}
	}
	if lastQuotient > 0 {
		curNode := &ListNode{
			Val: lastQuotient,
		}
		if tail == nil {
			tail = curNode
		} else {
			tail.Next = curNode
			tail = tail.Next
		}
	}

	return resList
}

func main() {
	var myList1 *ListNode
	var myList2 *ListNode
	addNode(&myList1, 9)
	addNode(&myList1, 9)
	addNode(&myList1, 9)
	addNode(&myList1, 9)
	addNode(&myList1, 9)
	addNode(&myList1, 9)
	addNode(&myList1, 9)
	addNode(&myList2, 9)
	addNode(&myList2, 9)
	addNode(&myList2, 9)
	addNode(&myList2, 9)

	res := addTwoNumbers(myList1, myList2)
	Traverse(res)
}
