package main

import (
	"fmt"
	"strings"
)

// Definire una funzione WordCount che data una stringa conta le occorrenze delle
// "parole" che la compongono, considerando gli spazi come separatori.
// Suggerimento: usare la funzione strings.Fields
func WordCount(s string) map[string]int {
	miseria := make(map[string]int)
	miserere := strings.Fields(s)
	for _, c := range miserere {
		miseria[c]++
	}
	return miseria
}

func main() {
	s := "aa aa aa aa aa aa aa bb bb bb bb bb bb bb bb cc cc cc cc cc cc cc cc sisisi sisisi sisisi si"
	var m map[string]int
	m = WordCount(s)
	for k, v := range m {
		fmt.Printf("%s\t%d\n", k, v)
	}
}
