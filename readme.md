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
go test --bench='(?i)bubble|gnome|heap|count'
```

### Sample results

After running

```bash
go test --bench='(?i)bubble|gnome|heap|count'
```

on my machine, I got the following results:

```bash
goos: linux
goarch: amd64
pkg: github.com/comparing_sorting_algorithms
BenchmarkBubbleSort/Permutations_of_1000_;_TOTAL_=_1000,_MAX_=_999_-8         	            308304	      3507 ns/op
BenchmarkBubbleSort/Permutations_of_10000_;_TOTAL_=_10000,_MAX_=_9999_-8      	             35311	     33358 ns/op
BenchmarkBubbleSort/ASC_Range_of_1000_elements_;_TOTAL_=_1000,_MAX_=_1000_-8  	            348283	      3421 ns/op
BenchmarkBubbleSort/ASC_Range_of_10000_elements_;_TOTAL_=_10000,_MAX_=_10000_-8         	   34801	     33768 ns/op
BenchmarkBubbleSort/DESC_Range_of_1000_elements_;_TOTAL_=_1000,_MAX_=_1000_-8           	  307598	      3432 ns/op
BenchmarkBubbleSort/DESC_Range_of_10000_elements_;_TOTAL_=_10000,_MAX_=_10000_-8        	   33542	     33567 ns/op
BenchmarkBubbleSort/Array_with_identical_elements:1_;_TOTAL_=_1000,_MAX_=_1_-8          	  321922	      3406 ns/op
BenchmarkBubbleSort/Array_with_identical_elements_1_;_TOTAL_=_10000,_MAX_=_1_-8         	   34862	     35161 ns/op

BenchmarkCountSort/Permutations_of_1000_;_TOTAL_=_1000,_MAX_=_999_-8                    	   43562	     28852 ns/op
BenchmarkCountSort/Permutations_of_10000_;_TOTAL_=_10000,_MAX_=_9999_-8                 	    8587	    175322 ns/op
BenchmarkCountSort/ASC_Range_of_1000_elements_;_TOTAL_=_1000,_MAX_=_1000_-8             	   39712	     31856 ns/op
BenchmarkCountSort/ASC_Range_of_10000_elements_;_TOTAL_=_10000,_MAX_=_10000_-8          	    7617	    164478 ns/op
BenchmarkCountSort/DESC_Range_of_1000_elements_;_TOTAL_=_1000,_MAX_=_1000_-8            	   39574	     29272 ns/op
BenchmarkCountSort/DESC_Range_of_10000_elements_;_TOTAL_=_10000,_MAX_=_10000_-8         	    7248	    171165 ns/op
BenchmarkCountSort/Array_with_identical_elements:1_;_TOTAL_=_1000,_MAX_=_1_-8           	   42853	     29379 ns/op
BenchmarkCountSort/Array_with_identical_elements_1_;_TOTAL_=_10000,_MAX_=_1_-8          	    7404	    156185 ns/op

BenchmarkHeapSort/Permutations_of_1000_;_TOTAL_=_1000,_MAX_=_999_-8                     	   17520	     66407 ns/op
BenchmarkHeapSort/Permutations_of_10000_;_TOTAL_=_10000,_MAX_=_9999_-8                  	    1318	    843976 ns/op
BenchmarkHeapSort/ASC_Range_of_1000_elements_;_TOTAL_=_1000,_MAX_=_1000_-8              	   16815	     67152 ns/op
BenchmarkHeapSort/ASC_Range_of_10000_elements_;_TOTAL_=_10000,_MAX_=_10000_-8           	    1344	    841160 ns/op
BenchmarkHeapSort/DESC_Range_of_1000_elements_;_TOTAL_=_1000,_MAX_=_1000_-8             	   16567	     69668 ns/op
BenchmarkHeapSort/DESC_Range_of_10000_elements_;_TOTAL_=_10000,_MAX_=_10000_-8          	    1404	    880674 ns/op
BenchmarkHeapSort/Array_with_identical_elements:1_;_TOTAL_=_1000,_MAX_=_1_-8            	  128612	      9611 ns/op
BenchmarkHeapSort/Array_with_identical_elements_1_;_TOTAL_=_10000,_MAX_=_1_-8           	   13208	     87779 ns/op

BenchmarkGnomeSort/Permutations_of_1000_;_TOTAL_=_1000,_MAX_=_999_-8                    	  321060	      3668 ns/op
BenchmarkGnomeSort/Permutations_of_10000_;_TOTAL_=_10000,_MAX_=_9999_-8                 	   30144	     36674 ns/op
BenchmarkGnomeSort/ASC_Range_of_1000_elements_;_TOTAL_=_1000,_MAX_=_1000_-8             	  289006	      3757 ns/op
BenchmarkGnomeSort/ASC_Range_of_10000_elements_;_TOTAL_=_10000,_MAX_=_10000_-8          	   32208	     36656 ns/op
BenchmarkGnomeSort/DESC_Range_of_1000_elements_;_TOTAL_=_1000,_MAX_=_1000_-8            	  296413	      3707 ns/op
BenchmarkGnomeSort/DESC_Range_of_10000_elements_;_TOTAL_=_10000,_MAX_=_10000_-8         	   32053	     36548 ns/op
BenchmarkGnomeSort/Array_with_identical_elements:1_;_TOTAL_=_1000,_MAX_=_1_-8           	  316794	      3678 ns/op
BenchmarkGnomeSort/Array_with_identical_elements_1_;_TOTAL_=_10000,_MAX_=_1_-8          	   31980	     36447 ns/op

PASS
ok  	github.com/comparing_sorting_algorithms	48.219s
```

*Note*: the second column indicates how many times the function was run within a second(by default). You can find more details [here](https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go).