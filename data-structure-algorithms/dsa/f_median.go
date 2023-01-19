package dsa

import "sort"

func MedianOfMedians(sliceList []int, k, r int) int {
	num := len(sliceList)
	if num < 10 {
		sort.Ints(sliceList)
		return sliceList[k-1]
	}
	med := (num + r - 1) / r

	medians := make([]int, med)

	for i := 0; i < med; i++ {
		v := (i * r) + r
		var arr []int
		if v >= num {
			arr = make([]int, len(sliceList[(i*r):]))
			copy(arr, sliceList[(i*r):])
		} else {
			arr = make([]int, r)
			copy(arr, sliceList[(i*r):v])
		}
		sort.Ints(arr)
		medians[i] = arr[len(arr)/2]
	}
	pivot := MedianOfMedians(medians, (len(medians)+1)/2, r)

	var leftSide, rightSide []int
	for i := range sliceList {
		if sliceList[i] < pivot {
			leftSide = append(leftSide, sliceList[i])
		} else if sliceList[i] > pivot {
			rightSide = append(rightSide, sliceList[i])
		}
	}

	switch {
	case k == (len(leftSide) + 1):
		return pivot
	case k <= len(leftSide):
		return MedianOfMedians(leftSide, k, r)
	default:
		return MedianOfMedians(rightSide, k-len(leftSide)-1, r)
	}
}
