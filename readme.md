# Comparing sorting algorithms

- [Comparing sorting algorithms](#comparing-sorting-algorithms)
  - [Prerequisites](#prerequisites)
  - [Sorting Algorithms](#sorting-algorithms)
  - [Results & specifications](#results--specifications)
    - [Running all the benchmarks](#running-all-the-benchmarks)
    - [Running benchmarks only for a specific algorithm](#running-benchmarks-only-for-a-specific-algorithm)
    - [Running benchmarks for specific algorithms](#running-benchmarks-for-specific-algorithms)
    - [Warning regarding `Stooge Sort`](#warning-regarding-stooge-sort)
    - [Sample results](#sample-results)
## Prerequisites

You'll need to have [Golang](https://golang.org/) installed on your machine.

---

## Sorting Algorithms

* [Bubble Sort](./algs/bubble.go)
* [Count Sort](./algs/count.go)
* [Heap Sort](./algs/heap.go)
* [Gnome Sort](./algs/gnome.go)
* [Stooge Sort](./algs/stooge.go)

---

## Results & specifications

### Running all the benchmarks

```bash
go test --bench=. 
```

### Running benchmarks only for a specific algorithm

```bash
go test --bench='(?i)<name>'
```

`--bench` takes in a regex and `(?i)` is used to indicate the **case-insensitive** modifier, so that you can simply type the name of the algorithm you'd want to see benchmarks for:

```bash
go test --bench='(?i)gnome'
go test --bench='(?i)heap'
```

### Running benchmarks for specific algorithms

```bash
go test --bench='(?i)gnome|heap'
```

### Warning regarding `Stooge Sort`

Because the `Stooge Sort` algorithm is **extremely** slow, I'd recommend you run the benchmarks for all the algorithms except this one. You can do so with the following command:

```bash
go test --bench='(?i)bubble|gnome|heap|count|native'
```

### Sample results

After running

```bash
go test --bench='(?i)bubble|gnome|heap|count|native'
```

on my machine, I got the following results:

```bash
goos: linux
goarch: amd64
pkg: github.com/comparing_sorting_algorithms
BenchmarkBubbleSort/Permutations_of_1000_;_TOTAL_=_1000,_MAX_=_999_-8         	            336240	      3682 ns/op
BenchmarkBubbleSort/Permutations_of_10000_;_TOTAL_=_10000,_MAX_=_9999_-8      	             32919	     37025 ns/op
BenchmarkBubbleSort/ASC_Range_of_1000_elements_;_TOTAL_=_1000,_MAX_=_1000_-8  	            309852	      3634 ns/op
BenchmarkBubbleSort/ASC_Range_of_10000_elements_;_TOTAL_=_10000,_MAX_=_10000_-8         	   32521	     34923 ns/op
BenchmarkBubbleSort/DESC_Range_of_1000_elements_;_TOTAL_=_1000,_MAX_=_1000_-8           	  289390	      3524 ns/op
BenchmarkBubbleSort/DESC_Range_of_10000_elements_;_TOTAL_=_10000,_MAX_=_10000_-8        	   33108	     34188 ns/op
BenchmarkBubbleSort/Array_with_identical_elements:1_;_TOTAL_=_1000,_MAX_=_1_-8          	  325699	      3568 ns/op
BenchmarkBubbleSort/Array_with_identical_elements_1_;_TOTAL_=_10000,_MAX_=_1_-8         	   32696	     35299 ns/op

BenchmarkCountSort/Permutations_of_1000_;_TOTAL_=_1000,_MAX_=_999_-8                    	   44132	     28481 ns/op
BenchmarkCountSort/Permutations_of_10000_;_TOTAL_=_10000,_MAX_=_9999_-8                 	    8372	    169169 ns/op
BenchmarkCountSort/ASC_Range_of_1000_elements_;_TOTAL_=_1000,_MAX_=_1000_-8             	   39830	     31804 ns/op
BenchmarkCountSort/ASC_Range_of_10000_elements_;_TOTAL_=_10000,_MAX_=_10000_-8          	    7312	    176724 ns/op
BenchmarkCountSort/DESC_Range_of_1000_elements_;_TOTAL_=_1000,_MAX_=_1000_-8            	   39724	     31132 ns/op
BenchmarkCountSort/DESC_Range_of_10000_elements_;_TOTAL_=_10000,_MAX_=_10000_-8         	    6600	    178965 ns/op
BenchmarkCountSort/Array_with_identical_elements:1_;_TOTAL_=_1000,_MAX_=_1_-8           	   43245	     29294 ns/op
BenchmarkCountSort/Array_with_identical_elements_1_;_TOTAL_=_10000,_MAX_=_1_-8          	    7166	    164025 ns/op

BenchmarkHeapSort/Permutations_of_1000_;_TOTAL_=_1000,_MAX_=_999_-8                     	   17745	     66997 ns/op
BenchmarkHeapSort/Permutations_of_10000_;_TOTAL_=_10000,_MAX_=_9999_-8                  	    1368	    827720 ns/op
BenchmarkHeapSort/ASC_Range_of_1000_elements_;_TOTAL_=_1000,_MAX_=_1000_-8              	   16636	     67112 ns/op
BenchmarkHeapSort/ASC_Range_of_10000_elements_;_TOTAL_=_10000,_MAX_=_10000_-8           	    1390	    804317 ns/op
BenchmarkHeapSort/DESC_Range_of_1000_elements_;_TOTAL_=_1000,_MAX_=_1000_-8             	   18225	     65635 ns/op
BenchmarkHeapSort/DESC_Range_of_10000_elements_;_TOTAL_=_10000,_MAX_=_10000_-8          	    1398	    801088 ns/op
BenchmarkHeapSort/Array_with_identical_elements:1_;_TOTAL_=_1000,_MAX_=_1_-8            	  126187	      9371 ns/op
BenchmarkHeapSort/Array_with_identical_elements_1_;_TOTAL_=_10000,_MAX_=_1_-8           	   12574	     90953 ns/op

BenchmarkGnomeSort/Permutations_of_1000_;_TOTAL_=_1000,_MAX_=_999_-8                    	  315808	      3717 ns/op
BenchmarkGnomeSort/Permutations_of_10000_;_TOTAL_=_10000,_MAX_=_9999_-8                 	   31620	     36664 ns/op
BenchmarkGnomeSort/ASC_Range_of_1000_elements_;_TOTAL_=_1000,_MAX_=_1000_-8             	  320732	      3688 ns/op
BenchmarkGnomeSort/ASC_Range_of_10000_elements_;_TOTAL_=_10000,_MAX_=_10000_-8          	   29839	     38849 ns/op
BenchmarkGnomeSort/DESC_Range_of_1000_elements_;_TOTAL_=_1000,_MAX_=_1000_-8            	  296325	      4398 ns/op
BenchmarkGnomeSort/DESC_Range_of_10000_elements_;_TOTAL_=_10000,_MAX_=_10000_-8         	   29257	     37843 ns/op
BenchmarkGnomeSort/Array_with_identical_elements:1_;_TOTAL_=_1000,_MAX_=_1_-8           	  303585	      3667 ns/op
BenchmarkGnomeSort/Array_with_identical_elements_1_;_TOTAL_=_10000,_MAX_=_1_-8          	   32768	     36711 ns/op

BenchmarkNativeSort/Permutations_of_1000_;_TOTAL_=_1000,_MAX_=_999_-8                   	   30529	     35769 ns/op
BenchmarkNativeSort/Permutations_of_10000_;_TOTAL_=_10000,_MAX_=_9999_-8                	    2275	    580263 ns/op
BenchmarkNativeSort/ASC_Range_of_1000_elements_;_TOTAL_=_1000,_MAX_=_1000_-8            	   29115	     42563 ns/op
BenchmarkNativeSort/ASC_Range_of_10000_elements_;_TOTAL_=_10000,_MAX_=_10000_-8         	    2318	    533351 ns/op
BenchmarkNativeSort/DESC_Range_of_1000_elements_;_TOTAL_=_1000,_MAX_=_1000_-8           	   33187	     35630 ns/op
BenchmarkNativeSort/DESC_Range_of_10000_elements_;_TOTAL_=_10000,_MAX_=_10000_-8        	    2536	    474735 ns/op
BenchmarkNativeSort/Array_with_identical_elements:1_;_TOTAL_=_1000,_MAX_=_1_-8          	  124982	      9980 ns/op
BenchmarkNativeSort/Array_with_identical_elements_1_;_TOTAL_=_10000,_MAX_=_1_-8         	   12296	     94886 ns/op

PASS
ok  	github.com/comparing_sorting_algorithms	59.602s
```

*Note*: the second column indicates how many times the function was run within a second(by default). You can find more details [here](https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go).