package main

func repeatedCharacter(s string) byte {
	b := []byte(s)
	check := map[byte]bool{}

	for i := 0; i < len(b); i++ {
		if _, ok := check[b[i]]; ok {
			return b[i]
		}

		check[b[i]] = true
	}

	return byte(0)
}
