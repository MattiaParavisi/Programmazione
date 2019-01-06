package main

import (
	"fmt"
	"strings"
)

func occorrenze(s, r string) int {
	var count int

	if len(s) == 0 {
		return 0
	}

	if strings.HasPrefix(s, string(r)) {
		count = 1
	}

	return count + occorrenze(s[1:], r)
}

func main() {
	var r string
	s := "cccc"
	fmt.Scan(&r)
	fmt.Println(occorrenze(s, r))
}
