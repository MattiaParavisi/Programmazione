package main

import (
	"fmt"
	"strings"
)

func main() {
	var a, b string
	fmt.Println("inserisci una stringa")
	fmt.Scan(&a)
	for _, c := range a {
		if strings.Index(b, string(c)) == -1 {
			b += string(c)
		}
	}
	fmt.Println(b)
}
