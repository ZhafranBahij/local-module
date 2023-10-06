package pengurangan

func Pengurangan(a int, b int) int {
	return a - b
}

func PenguranganBanyak(a int, values ...int) int {
	for _, v := range values {
		a -= v
	}

	return a
}
