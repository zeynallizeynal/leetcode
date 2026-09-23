package algorithms

//func MaxSubarraySum(nums []int, k int) int {
//	var sum int
//	for i := 0; i <= len(nums)-k; i++ {
//		var cur int
//		for j := i; j < k+i; j++ {
//			cur += nums[j]
//		}
//		if cur > sum {
//			sum = cur
//		}
//	}
//	return sum
//}

func MaxSubarraySum(nums []int, k int) int {
	var sum int
	for i := 0; i < k; i++ {
		sum += nums[i]
	}

	var cur = sum
	for i := 1; i <= len(nums)-k; i++ {
		cur -= nums[i-1]
		cur += nums[i+k-1]
		if cur > sum {
			sum = cur
		}
	}
	return sum
}

func FindMaxAverage(nums []int, k int) float64 {
	var res float64
	for i := 0; i < k; i++ {
		res += float64(nums[i])
	}

	var cur = res
	for i := 1; i <= len(nums)-k; i++ {
		cur -= float64(nums[i-1])
		cur += float64(nums[i+k-1])
		if cur > res {
			res = cur
		}
	}
	return res / float64(k)
}
