package algs

func GnomeSort(a []int) []int {
	idx := 0
	lenArr := len(a)

	for idx < lenArr {
		// this accounts for the case where we have `[10, 3, ...]`
		// in that case, the array will become `[3, 10, ...]`
		// and the index will become 0 and in the next iteration
		// we would get an error as we would want to access
		// an element which has a negative index
		if idx == 0 {
			idx++
			continue
		}

		if a[idx] < a[idx-1] {
			a[idx], a[idx-1] = a[idx-1], a[idx]
			// going one step backwards to that we check what happens
			// with the previous elements
			// who knows, this element my end up being the first one
			idx--
		} else {
			idx++
		}
	}

	return a
}
