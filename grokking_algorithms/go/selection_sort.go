package main

func findSmallest[T Numbers | string](lst []T) int {
	smallest := lst[0]
	smallestIndex := 0

	for i := 1; i < len(lst); i++ {
		if lst[i] < smallest {
			smallest = lst[i]
			smallestIndex = i
		}
	}

	return smallestIndex
}

func selectionSort[T Numbers | string](lst []T) []T {
	newLst := make([]T, len(lst))

	for i := range lst {
		smallestIndex := findSmallest(lst)
		newLst[i] = lst[smallestIndex]

		lst = append(lst[:smallestIndex], lst[smallestIndex+1:]...)
	}

	return newLst
}
