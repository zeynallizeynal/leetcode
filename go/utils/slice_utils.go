package utils

func SumSlice[T int64 | float64](s []T) T {
	var res T
	for _, v := range s {
		res += v
	}
	return res
}
