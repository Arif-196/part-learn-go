package b_sort

/**
Algoritma pengurutan cepat berada di bawah kelas algoritma bagi dan taklukkan, di mana kita memecah (membagi) masalah menjadi bagian yang lebih kecil yang lebih mudah dipecahkan (ditaklukkan). Dalam hal ini, array yang tidak disortir dipecah menjadi sub-array yang diurutkan sebagian, sampai semua elemen dalam daftar berada di posisi yang benar, saat daftar yang tidak disortir akan diurutkan.
*/

func SelectionSort(items []int) {
	var n = len(items)

	for i := 0; i < n; i++ {
		var minIdx = i
		for j := i; j < n; j++ {
			if items[j] < items[minIdx] {
				minIdx = j
			}
		}
		items[i], items[minIdx] = items[minIdx], items[i]
	}

}
