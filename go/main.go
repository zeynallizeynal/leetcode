package main

func main() {
	//list := NewLinkedListWithValue(10)
	list := NewLinkedListFromArray([]int{4, 13, 65, 7, 90, 54, 86, 0, 32})
	list.PushTail(3)
	list.PushHead(51)
	//list.PushHead(25)
	//list.DeleteHead()
	//list.DeleteTail()
	list.PrintFromHead()
	list.PrintFromTail()
}
