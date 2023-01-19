package b_sort

type Data []int32

func (dataList Data) Pancakesort() {
	for uns := len(dataList) - 1; uns > 0; uns-- {
		// find largest in unsorted range
		lx, lg := 0, dataList[0]
		for i := 1; i <= uns; i++ {
			if dataList[i] > lg {
				lx, lg = i, dataList[i]
			}
		}
		// move to final position in two flips
		dataList.Flip(lx)
		dataList.Flip(uns)
	}
}

func (dataList Data) Flip(r int) {
	for l := 0; l < r; l, r = l+1, r-1 {
		dataList[l], dataList[r] = dataList[r], dataList[l]
	}
}
