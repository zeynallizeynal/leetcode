package main

import (
	"fmt"
)

type ListNode struct {
	Value int
	Next  *ListNode
	Prev  *ListNode
}

type LinkedList struct {
	Head *ListNode
	Tail *ListNode
}

func NewLinkedListWithValue(val int) *LinkedList {
	var head = &ListNode{
		Value: val,
	}
	tail := head
	return &LinkedList{
		Head: head,
		Tail: tail,
	}
}

func NewLinkedListFromArray(arr []int) *LinkedList {
	if len(arr) == 0 {
		return nil
	}
	list := NewLinkedListWithValue(arr[0])
	head := list.Head

	for i := 1; i < len(arr); i++ {
		newNode := &ListNode{
			Value: arr[i],
		}
		head.Next = newNode
		newNode.Prev = head
		head = head.Next
	}
	list.Tail = head
	return list
}

func (l *LinkedList) PushTail(val int) {
	temp := l.Head
	for temp.Next != nil {
		temp = temp.Next
	}
	newNode := &ListNode{Value: val}
	temp.Next = newNode
	newNode.Prev = temp
	l.Tail = newNode
}

func (l *LinkedList) PushHead(val int) {
	temp := l.Head
	newNode := &ListNode{
		Value: val,
	}
	newNode.Next = temp
	temp.Prev = newNode
	l.Head = newNode
}

func (l *LinkedList) DeleteHead() {
	if l.Head == nil {
		return
	}
	temp := l.Head

	if temp.Next != nil {
		temp.Next.Prev = nil
		l.Head = temp.Next
		temp.Next = nil
	} else {
		l.Tail = nil
		l.Head = nil
	}
}

func (l *LinkedList) DeleteTail() {
	if l.Tail == nil {
		return
	}
	temp := l.Tail

	if temp.Prev != nil {
		temp.Prev.Next = nil
		l.Tail = temp.Prev
		temp.Prev = nil
	} else {
		l.Tail = nil
		l.Head = nil
	}
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
		fmt.Printf("%d -> ", temp.Value)
		temp = temp.Prev
	}
	fmt.Println("nil")
}
