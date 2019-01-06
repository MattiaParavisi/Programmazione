package main

import "fmt"

func main() {
	var s string

	mappa := make(map[rune]int)

	fmt.Scan(&s)

	for _, c := range s {
		if 1 != mappa[c] {
			mappa[c]++
		} else {

		}
	}

	for c := range mappa {
		fmt.Printf("%c", c)
	}
}
