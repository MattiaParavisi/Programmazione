package main

import "fmt"

func IsEqual(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}

func main() {
	var s, anagramma string
	fmt.Scan(&s, &anagramma)
	ok := IsEqual(s, anagramma)
	if ok {
		fmt.Println("Uguali")
	} else {
		fmt.Println("Non uguali")
	}
}
