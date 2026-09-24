package algorithms

func SmallestIndex(nums []int) int {
	for i := 0; i < len(nums); i++ {
		cur := nums[i]
		sum := 0
		for cur > 0 {
			sum += (cur % 10)
			cur /= 10
		}
		if sum == i {
			return i
		}
	}
	return -1
}
