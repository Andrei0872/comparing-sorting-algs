package algs

/*
- for an element `i`, the `2i+1` element is the left child and the `2i+2` element is the right child
- for an element `i`, its **parent** is the `(i-1)/2` element
- the index of the fist **non-leaf** element is `n/2 - 1`

A (min/max) heap must:
	* be a complete binary tree
	* each node must be greater(max heap)/smaller(min heap) than its children
*/

// fitting the node at the index `i` into the heap
func heapify(arr []int, lenArr, i int) {
	parentIdx := i
	leftChildIdx := 2*i + 1
	rightChildIdx := 2*i + 2

	var largest = parentIdx

	// the left child is bigger - we're going the `left` route
	if leftChildIdx < lenArr && arr[largest] < arr[leftChildIdx] {
		largest = leftChildIdx
	}

	// the right child is bigger - we're going the `right` route
	if rightChildIdx < lenArr && arr[largest] < arr[rightChildIdx] {
		largest = rightChildIdx
	}

	// if either the left child or right child is bigger
	if largest != parentIdx {
		arr[parentIdx], arr[largest] = arr[largest], arr[parentIdx]

		// repeating the process again, but this time the initial `i-node`'s children
		// are the ones of the largest(left or right) child found
		heapify(arr, lenArr, largest)
	}
}

func HeapSort(a []int) []int {
	lenArr := len(a)

	// the base idea is to construct the sub-trees first
	// in order to to that we'll start from the node of index `n/2 -1`
	// because this is **first node** which can't be a leaf
	// and we will `heapify` this subtree, which will have for sure only leaf nodes
	// then, we'll go backwords, to the other non-leaf nodes
	// this way, we're `heapifying` the sub-trees in a bottom-up fashion
	for i := (lenArr / 2) - 1; i >= 0; i-- {
		heapify(a, lenArr, i)
	}

	for lenArr > 0 {
		lenArr--

		lastElem := a[lenArr]
		firstElem := a[0]

		a[0] = lastElem
		a[lenArr] = firstElem

		// `heapifying` from root, where the root is the element found at the (new)last position
		// the previous root now lays at the end of the array - that's how the array essentially gets sorted
		heapify(a, lenArr, 0)
	}

	return a
}
