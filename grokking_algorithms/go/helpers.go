package main

func contains[K comparable](v K, arr []K) bool {
	for _, i := range arr {
		if i == v {
			return true
		}
	}
	return false
}

func isSeller(name string) bool {
	if name == "you" {
		return false
	}

	return name[len(name)-1] == 'm'
}
