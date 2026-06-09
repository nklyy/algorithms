package main

func largestAltitude(gain []int) int {
	prefixGain := []int{0}
	maxAltitude := prefixGain[0]

	for i := 0; i < len(gain); i++ {
		s := prefixGain[len(prefixGain)-1] + gain[i]

		if s > maxAltitude {
			maxAltitude = s
		}
		prefixGain = append(prefixGain, s)
	}

	return maxAltitude
}

// Alternative solution without using extra space for prefix sums
// func largestAltitude(gain []int) int {
// 	currentAltitude := 0
// 	maxAltitude := 0

// 	for _, g := range gain {
// 		currentAltitude += g
// 		if currentAltitude > maxAltitude {
// 			maxAltitude = currentAltitude
// 		}
// 	}

// 	return maxAltitude
// }
