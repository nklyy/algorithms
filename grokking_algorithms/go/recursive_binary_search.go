package main

func recursiveBinarySearchRecursive[T Numbers | string](lst []T, i T, low, high int) int {
	if low > high {
		return -1
	}

	mid := (low + high) / 2
	if lst[mid] == i {
		return mid
	}

	if lst[mid] < i {
		return recursiveBinarySearchRecursive(lst, i, mid+1, high)
	}
	return recursiveBinarySearchRecursive(lst, i, low, mid-1)
}
