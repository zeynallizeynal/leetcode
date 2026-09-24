package main

import "fmt"

type ListNode struct {
	Value interface{}
	Next  *ListNode
	Prev  *ListNode
}

type LinkedList struct {
	Head *ListNode
	Tail *ListNode
	Size int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func NewLinkedListFromArray(arr []int) *LinkedList {
	list := NewLinkedList()
	for _, v := range arr {
		list.PushTail(v)
	}
	return list
}

func (l *LinkedList) PushTail(val int) {
	newNode := &ListNode{Value: val}
	if l.Tail == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		newNode.Prev = l.Tail
		l.Tail.Next = newNode
		l.Tail = newNode
	}
	l.Size++
}

func (l *LinkedList) PushHead(val int) {
	newNode := &ListNode{Value: val}
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		newNode.Next = l.Head
		l.Head.Prev = newNode
		l.Head = newNode
	}
	l.Size++
}

func (l *LinkedList) DeleteHead() {
	if l.Head == nil {
		return
	}
	l.Head = l.Head.Next
	if l.Head != nil {
		l.Head.Prev = nil
	} else {
		l.Tail = nil
	}
	l.Size--
}

func (l *LinkedList) DeleteTail() {
	if l.Tail == nil {
		return
	}
	l.Tail = l.Tail.Prev
	if l.Tail != nil {
		l.Tail.Next = nil
	} else {
		l.Head = nil
	}
	l.Size--
}

func (l *LinkedList) GetTail() *ListNode {
	return l.Tail
}

func (l *LinkedList) GetHead() *ListNode {
	return l.Head
}

func (l *LinkedList) PrintFromHead() {
	temp := l.Head
	for temp != nil {
		fmt.Printf("%d -> ", temp.Value)
		temp = temp.Next
	}
	fmt.Println("nil")
}

func (l *LinkedList) PrintFromTail() {
	temp := l.Tail
	for temp != nil {
		fmt.Printf("%d <- ", temp.Value)
		temp = temp.Prev
	}
	fmt.Println("nil")
}

//func MergeSortedLists(l1, l2 *ListNode) *ListNode {
//	res := &ListNode{}
//	h1 := l1
//	h2 := l2
//}
