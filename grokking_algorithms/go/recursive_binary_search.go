package main

func recursiveBinarySearchRecursive[T Numbers | string](list []T, i T, low, high int) int {
	if low > high {
		return -1
	}

	mid := (low + high) / 2
	if list[mid] == i {
		return mid
	}

	if list[mid] < i {
		return recursiveBinarySearchRecursive(list, i, mid+1, high)
	}
	return recursiveBinarySearchRecursive(list, i, low, mid-1)
}
