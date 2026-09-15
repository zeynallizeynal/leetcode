package main

import "fmt"

func main() {
	//list := NewLinkedListWithValue(10)
	//list := NewLinkedListFromArray([]int{4, 13, 65, 7, 90, 54, 86, 0, 32})
	//list.PushTail(3)
	//list.PushHead(51)
	//list.PushHead(25)
	//list.DeleteHead()
	//list.DeleteTail()
	//list.PrintFromHead()
	//list.PrintFromTail()

	circArr := NewCircularArray(5)
	circArr.AddBack(5)
	circArr.AddBack(23)
	fmt.Println(*circArr.GetBack())
	fmt.Println(*circArr.GetFront())
}
