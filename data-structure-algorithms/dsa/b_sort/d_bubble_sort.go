package b_sort

import (
	"math/rand"
	"time"
)

/**
Diberikan daftar yang tidak terurut, kami membandingkan elemen yang berdekatan dalam daftar, setiap kali, menempatkan dalam urutan besarnya yang benar, hanya dua elemen. Algoritme bergantung pada prosedur swap. Mengetahui berapa kali bertukar adalah penting saat menerapkan algoritme pengurutan gelembung. Untuk mengurutkan daftar angka seperti [3, 2, 1], kita perlu menukar elemen maksimal dua kali. Ini sama dengan panjang daftar dikurangi 1
*/

func GeneratSlice(size int) []int {
	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func BubbleSort(items []int) {
	var (
		n      = len(items)
		sorted = false
	)

	for !sorted {
		swapped := false
		for i := 0; i < n-1; i++ {
			if items[i] > items[i+1] {
				items[i+1], items[i] = items[i], items[i+1]
				swapped = true
			}
		}
		if !swapped {
			sorted = true
		}
		n = n - 1
	}
}
