package algs

func stoogeSort(a []int, firstIdx, lastIdx int) {
	if firstIdx >= lastIdx {
		return
	}

	if a[firstIdx] > a[lastIdx] {
		a[firstIdx], a[lastIdx] = a[lastIdx], a[firstIdx]
	}

	sliceLen := lastIdx - firstIdx + 1

	if !(sliceLen > 2) {
		return
	}

	fraction := sliceLen / 3

	// first 2/3 elements
	stoogeSort(a, firstIdx, lastIdx-fraction)
	// last 2/3 elements
	stoogeSort(a, firstIdx+fraction, lastIdx)

	// making sure the elements stay sorted after sorting the last 2/3 elements
	stoogeSort(a, firstIdx, lastIdx-fraction)
}

func StoogeSort(a []int) []int {
	stoogeSort(a, 0, len(a)-1)

	return a
}
