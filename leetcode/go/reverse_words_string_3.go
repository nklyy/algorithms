package main

// func reverseWords(s string) string {
// 	var res []string

// 	for _, v := range strings.Split(s, " ") {
// 		l := 0
// 		r := len(v) - 1
// 		rW := []rune(v)

// 		for l < r {
// 			rW[l], rW[r] = rW[r], rW[l]
// 			l++
// 			r--
// 		}

// 		res = append(res, string(rW))
// 	}

// 	return strings.Join(res, " ")
// }

// More fast solution without using extra space for words array and runes array.
func reverseWords(s string) string {
	b := []byte(s)
	start := 0

	for i := 0; i <= len(b); i++ {
		if i == len(b) || b[i] == ' ' {
			l, r := start, i-1

			for l < r {
				b[l], b[r] = b[r], b[l]
				l++
				r--
			}

			start = i + 1
		}
	}

	return string(b)
}
