package penjumlahan

func Penjumlahan(a int, b int) int {
	sum := a + b
	return sum
}

func PenjumlahanBanyak(values ...int) int {
	sum := 0
	for _, v := range values {
		sum += v
	}
	return sum
}
