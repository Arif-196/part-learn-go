package b_sort

import "math/rand"

/**
Algoritma pengurutan cepat berada di bawah kelas algoritma bagi dan taklukkan, di mana kita memecah (membagi) masalah menjadi bagian yang lebih kecil yang lebih mudah dipecahkan (ditaklukkan). Dalam hal ini, array yang tidak disortir dipecah menjadi sub-array yang diurutkan sebagian, sampai semua elemen dalam daftar berada di posisi yang benar, saat daftar yang tidak disortir akan diurutkan.
*/
func QuickSort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := rand.Int() % len(a)

	a[pivot], a[right] = a[right], a[pivot]

	for i, _ := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}
	a[left], a[right] = a[right], a[left]

	QuickSort(a[:left])
	QuickSort(a[left+1:])

	return a
}
