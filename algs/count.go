package algs

import "fmt"

const MAX_ACCEPTED_NUMBER = 1e4 + 1

func CountSort(a []int) []int {
	freq := make([]int, MAX_ACCEPTED_NUMBER)
	res := []int{}

	for _, elem := range a {
		if elem > int(MAX_ACCEPTED_NUMBER) {
			panic(fmt.Sprintf("The current Count Sort algorithm does not support numbers greater than %d !\n", int(MAX_ACCEPTED_NUMBER)))
		}

		freq[elem]++
	}

	for i := 0; i < int(MAX_ACCEPTED_NUMBER); i++ {
		if freq[i] == 0 {
			continue
		}

		for j := 0; j < freq[i]; j++ {
			res = append(res, i)
		}
	}

	return res
}
