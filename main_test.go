package main

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/comparing_sorting_algorithms/algs"
)

// TODO: use native sort as well

const (
	N = 1e4
)

// 0...N
func genPermOfN(N int) []int {
	rand.Seed(time.Now().Unix())

	return rand.Perm(N)
}

func genRange(min, max int, reverse bool) []int {
	res := make([]int, max-min+1)

	for i := range res {
		var val int

		if !reverse {
			val = min + i
		} else {
			val = max - i
		}

		res[i] = val
	}

	return res
}

func genArrayWithIdenticalElements(N, elem int) []int {
	res := make([]int, N)

	for i := range res {
		res[i] = elem
	}

	return res
}

var tableTests = []struct {
	totalNums   int
	maxNum      int
	list        []int
	title       string
	description string
}{
	{int(N / 10), (int(N) / 10) - 1, genPermOfN(N / 10), fmt.Sprintf("TOTAL = %d, MAX = %d\n", int(N)/10, (int(N)/10)-1), fmt.Sprintf("Permutations of %d\n", int(N/10))},
	{int(N), int(N) - 1, genPermOfN(N), fmt.Sprintf("TOTAL = %d, MAX = %d\n", int(N), int(N)-1), fmt.Sprintf("Permutations of %d\n", int(N))},

	{int(N / 10), (int(N) / 10), genRange(1, N/10, false), fmt.Sprintf("TOTAL = %d, MAX = %d\n", int(N)/10, (int(N) / 10)), fmt.Sprintf("ASC Range of %d elements\n", int(N/10))},
	{int(N), int(N) - 1, genRange(1, N, false), fmt.Sprintf("TOTAL = %d, MAX = %d\n", int(N), int(N)), fmt.Sprintf("ASC Range of %d elements\n", int(N))},

	{int(N / 10), (int(N) / 10), genRange(1, N/10, true), fmt.Sprintf("TOTAL = %d, MAX = %d\n", int(N)/10, (int(N) / 10)), fmt.Sprintf("DESC Range of %d elements\n", int(N/10))},
	{int(N), int(N) - 1, genRange(1, N, true), fmt.Sprintf("TOTAL = %d, MAX = %d\n", int(N), int(N)), fmt.Sprintf("DESC Range of %d elements\n", int(N))},

	{int(N / 10), (int(N) / 10) - 1, genArrayWithIdenticalElements(N/10, 1), fmt.Sprintf("TOTAL = %d, MAX = %d\n", int(N)/10, 1), fmt.Sprintf("Array with identical elements:%d\n", 1)},
	{int(N), int(N) - 1, genArrayWithIdenticalElements(N, 1), fmt.Sprintf("TOTAL = %d, MAX = %d\n", int(N), 1), fmt.Sprintf("Array with identical elements %d\n", 1)},
}

func BenchmarkBubble(t *testing.B) {
	for _, test := range tableTests {
		t.Run(fmt.Sprintf("%s; %s", test.description, test.title), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				isSorted := sort.IntsAreSorted(algs.BubbleSort(test.list))
				if !isSorted {
					b.Errorf("The current algorithm did not properly sorted the elements!")
				}
			}
		})
	}

	fmt.Println()
}
