package main

func factorial(x uint) uint {
	if x == 1 {
		return 1
	}

	return x * factorial(x-1)
}

func factorialLoop(x int) int {
	if x == 0 || x == 1 {
		return 1
	}

	r := x
	for i := x; i > 1; i-- {
		r = r * (i - 1)
	}

	// r := 1
	// for i := 1; i <= x; i++ {
	// 	r = r * i
	// }

	return r
}
