package main

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

func BinarySearch[T Signed | string](list []T, i T) int {
	low := 0
	high := len(list) - 1

	for low <= high {
		mid := (low + high) / 2

		if list[mid] == i {
			return mid
		}

		if list[mid] < i {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}
