package main

func binarySearch[T Numbers | string](lst []T, i T) int {
	low := 0
	high := len(lst) - 1

	for low <= high {
		mid := (low + high) / 2

		if lst[mid] == i {
			return mid
		}

		if lst[mid] < i {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}
