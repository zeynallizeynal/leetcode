package main

func main() {
	// list := NewLinkedList()
	// list := NewLinkedListFromArray([]int{4, 13, 65, 7, 90, 54, 86, 0, 32})
	// list.PushTail(3)
	// list.PushHead(51)
	// list.PushHead(25)
	// list.DeleteHead()
	// list.DeleteTail()
	// list.PrintFromHead()
	// list.PrintFromTail()

	// circArr := NewCircularArray(5)
	// circArr.AddBack(5)
	// circArr.AddBack(23)
	// fmt.Println(*circArr.GetBack())
	// fmt.Println(*circArr.GetFront())

	// stack := NewLinkedStack()
	// fmt.Println(stack.Size())

	// stack.Push(10)
	// stack.Push(146)
	// fmt.Println(stack.Size())
	// fmt.Println(stack.Peek())
	// stack.Pop()
	// fmt.Println(stack.Size())
	// fmt.Println(stack.Peek())
	// stack.Print()

	queue := NewLinkedQueue()
	queue.Push(10)
	queue.Push(8)
	queue.Print()
}
