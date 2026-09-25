package main

import (
	"fmt"

	"github.com/zeynallizeynal/leetcode/algorithms"
)

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

	//circArr := NewCircularArray(5)
	//circArr.AddBack(5)
	//circArr.AddBack(23)
	//fmt.Println(*circArr.GetBack())
	//fmt.Println(*circArr.GetFront())

	//table := NewHashTable[string, int](1)
	//table.Put("zeynal", 31)
	//table.Put("aydan", 90)
	//fmt.Println(table.Get("aydan"))

	//trie := NewTrie()
	//trie.Insert("hello")
	//trie.Insert("help")
	//fmt.Println(trie.StartsWith("hel"))

	// algorithm techniques
	// nums := []int{1, 12, -5, -6, 50, 3}
	// fmt.Println(algorithms.FindMaxAverage(nums, 4))
	// fmt.Println(algorithms.ContainsDuplicate([]int{1, 2, 3, 4, 1}))
	// fmt.Println(algorithms.ContainsNearbyDuplicate([]int{1, 0, 1, 1}, 1))

	// fmt.Println(algorithms.SmallestIndex([]int{1, 3, 2}))

	// fmt.Println(algorithms.MaxSubarrayProduct([]int{1, 4, 1, 6, -3, 3, -5, 2, 26}, 4))
	//fmt.Println(algorithms.TargetSubarraySum([]int{2, 3, 2, 2, 3, 1, 3, 8, 5, 0, 2, 4}, 7, 3))

	fmt.Println(algorithms.StringNaiveMatch("mississippi", "issipi"))
}
