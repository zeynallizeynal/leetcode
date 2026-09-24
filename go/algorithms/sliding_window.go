package algorithms

import "math"

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

func ContainsDuplicate(nums []int) bool {
	m := make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			return true
		}
		m[nums[i]] = struct{}{}
	}
	return false
}

func ContainsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		n, ok := m[nums[i]]
		if ok && math.Abs(float64(i-n)) <= float64(k) {
			return true
		}
		m[nums[i]] = i
	}
	return false
}
