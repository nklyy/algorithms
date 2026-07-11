package main

func checkIfPangram(sentence string) bool {
	sentenceByte := []byte(sentence)
	checkMap := map[byte]bool{}

	for i := 0; i < len(sentenceByte); i++ {
		checkMap[sentenceByte[i]] = true
	}

	return len(checkMap) == 26
}
