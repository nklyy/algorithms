package main

// Finds the smallest value in an array
func findSmallest(arr []int) int {
	smallest := arr[0]
	smallestIndex := 0

	for i := 1; i < len(arr); i++ {
		if arr[i] < smallest {
			smallest = arr[i]
			smallestIndex = i
		}
	}

	return smallestIndex
}

// Sort array
func selectionSort(arr []int) []int {
	newArr := make([]int, len(arr))

	for i := range arr {
		smallestIndex := findSmallest(arr)
		newArr[i] = arr[smallestIndex]

		arr = append(arr[:smallestIndex], arr[smallestIndex+1:]...)
	}

	return newArr
}
