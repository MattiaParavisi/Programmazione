package main

import "fmt"

func isAnagr(x string, y string) bool {

	map1 := make(map[rune]int)
	map2 := make(map[rune]int)

	if x == "" && y == "" {
		return true
	}

	for _, i := range x {
		map1[i]++
	}

	for _, i := range y {
		map2[i]++
	}

	for k, v := range map1 {
		if map2[k] != v {
			return false
		}
	}

	for k, v := range map2 {
		if map1[k] != v {
			return false
		}
	}

	return true
}

func main() {
	var string1, string2 string
	fmt.Scan(&string1, &string2)
	if isAnagr(string1, string2) {
		fmt.Println("palindromi")
	} else {
		fmt.Println("non palindrome")
	}
}
