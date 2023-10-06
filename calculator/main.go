package main

import (
	"calculator/pengurangan"
	"calculator/penjumlahan"
	"fmt"
)

func main() {
	fmt.Println("Melakukan Perhitungan")
	fmt.Println(penjumlahan.Penjumlahan(5, 2))
	fmt.Println(penjumlahan.PenjumlahanBanyak(1, 2, 3, 4, 5))
	fmt.Println(pengurangan.Pengurangan(5, 2))
	fmt.Println(pengurangan.PenguranganBanyak(1, 2, 3, 4, 5))
}
