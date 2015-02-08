package natural

func Sum(x []int64) int64 {
	sum := int64(0)
	for _, y := range x {
		sum += y
	}
	return sum
}

func Min(args ...int64) int64 {
	min := args[0]
	for _, v := range args[1:] {
		if min > v {
			min = v
		}
	}
	return min
}

func Max(args ...int64) int64 {
	max := args[0]
	for _, v := range args[1:] {
		if max < v {
			max = v
		}
	}
	return max
}
