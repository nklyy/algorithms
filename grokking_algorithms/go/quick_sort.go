package main

func quickSort[T Numbers | string](arr []T) []T {
	if len(arr) < 2 {
		return arr
	}

	if len(arr) == 2 {
		if arr[0] > arr[1] {
			arr[0], arr[1] = arr[1], arr[0]
		}

		return arr
	}

	pivot := arr[0]
	less := []T{}
	greater := []T{}

	for _, v := range arr[1:] {
		if v < pivot {
			less = append(less, v)
		} else {
			greater = append(greater, v)
		}
	}

	result := []T{}
	result = append(result, quickSort(less)...)
	result = append(result, pivot)
	result = append(result, quickSort(greater)...)

	return result
}

// func quickSort[T Numbers | string](arr []T) []T {
// 	if len(arr) < 2 {
// 		return arr
// 	} else {
// 		pivot := arr[0]
// 		less := []T{}
// 		greater := []T{}

// 		for _, v := range arr[1:] {
// 			if v < pivot {
// 				less = append(less, v)
// 			} else {
// 				greater = append(greater, v)
// 			}
// 		}

// 		result := []T{}
// 		result = append(result, quickSort(less)...)
// 		result = append(result, pivot)
// 		result = append(result, quickSort(greater)...)

// 		return result
// 	}
// }
