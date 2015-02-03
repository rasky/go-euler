package natural

func Sum(x []int64) int64 {
	sum := int64(0)
	for _, y := range x {
		sum += y
	}
	return sum
}
