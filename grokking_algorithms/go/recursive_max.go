package main

func recursiveMax(lst []int) int {
	if len(lst) == 2 {
		if lst[0] > lst[1] {
			return lst[0]
		}
		return lst[1]
	}

	max := recursiveMax(lst[1:])
	if lst[0] > max {
		return lst[0]
	}

	return max
}
