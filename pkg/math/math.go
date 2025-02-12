package math

func PowInt64(x, y int64) int64 {
	res := int64(1)
	for i := int64(0); i < y; i++ {
		res *= int64(x)
	}

	return res
}
