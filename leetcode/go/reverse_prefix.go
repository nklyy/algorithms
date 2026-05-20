package main

func reversePrefix(word string, ch byte) string {
	b := []byte(word)

	for i := 0; i < len(b); i++ {
		if b[i] == ch {
			l, r := 0, i
			for l < r {
				b[l], b[r] = b[r], b[l]
				l++
				r--
			}

			return string(b)
		}
	}

	return word
}

// Alternative solution using strings.IndexByte
// func reversePrefix(word string, ch byte) string {
//     i := strings.IndexByte(word, ch)
//     if i == -1 {
//         return word
//     }
//     b := []byte(word)
//     for l, r := 0, i; l < r; l, r = l+1, r-1 {
//         b[l], b[r] = b[r], b[l]
//     }
//     return string(b)
// }
