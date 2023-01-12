package main

func quickSort[T Numbers | string](lst []T) []T {
	if len(lst) < 2 {
		return lst
	}

	if len(lst) == 2 {
		if lst[0] > lst[1] {
			lst[0], lst[1] = lst[1], lst[0]
		}

		return lst
	}

	pivot := lst[0]
	less := []T{}
	greater := []T{}

	for _, v := range lst[1:] {
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

// func quickSort[T Numbers | string](lst []T) []T {
// 	if len(lst) < 2 {
// 		return lst
// 	} else {
// 		pivot := lst[0]
// 		less := []T{}
// 		greater := []T{}

// 		for _, v := range lst[1:] {
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
