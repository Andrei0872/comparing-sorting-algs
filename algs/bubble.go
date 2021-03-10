package algs

func BubbleSort(arr []int) []int {
	isSorted := false
	arrLen := len(arr)

	for !isSorted {
		i := 0
		isSorted = true

		for ; i < arrLen-1; i++ {
			if arr[i] > arr[i+1] {
				temp := arr[i+1]
				arr[i+1] = arr[i]
				arr[i] = temp

				isSorted = false
			}
		}
	}

	return arr
}
