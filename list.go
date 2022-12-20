package main

import (
	"fmt"
)

type Node struct {
	value int
	next  *Node
}

type List struct {
	head *Node
}

func (p *List) Append(value int) {
	// 如果head为空，直接给头部
	if p.head == nil {
		p.head = &Node{
			value: value,
		}
		return
	}
	curNode := p.head
	for ; curNode.next != nil; curNode = curNode.next {
	}
	curNode.next = &Node{
		value: value,
	}
}

func (p *List) Traverse() {
	for curNode := p.head; curNode != nil; curNode = curNode.next {
		fmt.Printf("curNode' value is %d\n", curNode.value)
	}
	fmt.Println("traverse over!")
}

func (p *List) Reverse() {
	// 边界条件直接返回
	if p.head == nil || p.head.next == nil {
		return
	}
	curNode := p.head
	nextNode := p.head.next
	var nextNextNode *Node
	for nextNode != nil {
		nextNextNode = nextNode.next
		if curNode == p.head {
			curNode.next = nil
		}
		nextNode.next = curNode
		curNode = nextNode
		nextNode = nextNextNode
	}
	p.head = curNode
}

func (p *List) ReverseByGroupOfNum(num int) {
	// param check
	if num <= 0 {
		return
	}
	if p.head == nil {
		return
	}
	curNode := p.head
	var (
		groupHead     *Node
		groupTail     *Node
		nextGroupHead *Node
		LastGroupTail *Node
	)
	groupHead = curNode
	cnt := 1
	var flag bool
	for curNode.next != nil {
		if cnt%num == 0 {
			groupTail = curNode
			nextGroupHead = curNode.next
			groupTail.next = nil
			tmpGroupList := &List{
				head: groupHead,
			}
			tmpGroupList.Reverse()
			if LastGroupTail == nil {
				flag = true
				p.head = tmpGroupList.head
			} else {
				LastGroupTail.next = groupHead
			}
			LastGroupTail = groupHead
			groupHead = nextGroupHead
			curNode = nextGroupHead
		} else {
			curNode = curNode.next
		}
		cnt++
	}
	groupTail = curNode
	nextGroupHead = curNode.next
	groupTail.next = nil
	tmpGroupList := &List{
		head: groupHead,
	}
	tmpGroupList.Reverse()
	if LastGroupTail != nil {
		LastGroupTail.next = tmpGroupList.head
	}
	LastGroupTail = groupHead
	if !flag {
		p.head = tmpGroupList.head
	}
}

func main() {
	myList := List{}
	myList.Append(1)
	myList.Append(3)
	myList.Append(2)
	myList.Append(4)
	myList.ReverseByGroupOfNum(1)
	myList.Traverse()
}
