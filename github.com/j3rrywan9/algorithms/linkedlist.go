package main

import (
	"fmt"
)

type ListNode struct {
	val int
	next *ListNode
}

func (this *ListNode) CreateLinkedList(length int) {
	current := this
	for i := 2; i <= length; i++ {
		pListNode := &ListNode{i, nil}
		current.next = pListNode
		current = pListNode
	}
}

func (this *ListNode) PrintLinkedList() {
	current := this
	for {
		fmt.Printf("%d", current.val)
		if current.next != nil {
			fmt.Print("->")
			current = current.next
		} else {
			break
		}
	}
	fmt.Println()
}

// LCOJ No. 206
func (this *ListNode) ReverseLinkedList(pListNode *ListNode) *ListNode {
	current := pListNode
	var head *ListNode = nil
	for {
		if current == nil {
			break
		}
		temp := current.next
		current.next = head
		head = current
		current = temp
	}
	return head
}

// LCOJ No. 237
func (this *ListNode) DeleteNode(pListNode *ListNode) {
	if pListNode != nil && pListNode.next != nil {
		pListNode.val = pListNode.next.val
		pListNode.next = pListNode.next.next
	}
}

func main() {
	mylistnode := &ListNode{1, nil}
	mylistnode.CreateLinkedList(10)
	mylistnode.DeleteNode(mylistnode.next.next)
	mylistnode.PrintLinkedList()
	mylistnode2 := mylistnode.ReverseLinkedList(mylistnode)
	mylistnode2.PrintLinkedList()
}

