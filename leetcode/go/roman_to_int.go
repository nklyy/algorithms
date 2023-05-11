package main

var symbols = map[string]int{
	"I":  1,
	"IV": 4,
	"V":  5,
	"IX": 9,
	"X":  10,
	"XL": 40,
	"L":  50,
	"XC": 90,
	"C":  100,
	"CD": 400,
	"D":  500,
	"CM": 900,
	"M":  1000,
}

func romanToInt(s string) int {
	res := 0
	l := 0

	for l < len(s) {
		if l+1 >= len(s) {
			res += symbols[string(s[l])]
			return res
		}

		if v, ok := symbols[string(s[l])+string(s[l+1])]; ok {
			res += v
			l += 2
		} else {
			res += symbols[string(s[l])]
			l++
		}
	}

	return res
}

// Other version
// var symbols = map[string]int{
// 	"I": 1,
// 	"V": 5,
// 	"X": 10,
// 	"L": 50,
// 	"C": 100,
// 	"D": 500,
// 	"M": 1000,
// }

// func romanToInt(s string) int {
// 	res := 0
// 	l := len(s) - 1
// 	for i := 0; i < l; i++ {
// 		if symbols[string(s[i])] < symbols[string(s[i+1])] {
// 			res -= symbols[string(s[i])]
// 			fmt.Println(res)
// 		} else {
// 			res += symbols[string(s[i])]
// 		}
// 	}
// 	res += symbols[string(s[l])]
// 	return res
// }
