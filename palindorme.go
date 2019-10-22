package main

import "fmt"

func IterIsPalin(s string) bool {
	if len(s) <= 1 {
		return true
	}
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1] {
			return false
		}
	}
	return true
}

func main() {
	var s string
	fmt.Scan(&s)
	ok := IterIsPalin(s)
	if ok {
		fmt.Println("palindorma")
	} else {
		fmt.Println("non palindorma")
	}
}
