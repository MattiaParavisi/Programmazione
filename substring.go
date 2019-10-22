package main

import (
	"fmt"
	"strings"
)

func contains(s string, parola string) int {
	count := 0

	if len(parola) > len(s) {
		return 0
	}

	if strings.HasPrefix(s, parola) {

		count++

	}

	return count + contains(s[1:], parola)
}

func main() {
	s := "ciaociaocoiaciaoiaciocia"
	fmt.Println(s)
	fmt.Print("stringa da cercare: ")
	var sub string
	fmt.Scanf("%s", &sub)
	fmt.Printf("%s sta %d volte in %s\n", sub, contains(s, sub), s)
}

// calcola, in modo ricorsivo, quante volte sub Ã¨ contenuta in s
