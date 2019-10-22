package main

import (
	"fmt"
	"os"
	"strings"
)

func contaOccorrenze(s string) [26]int {
	var occorrenza [26]int
	s = strings.ToLower(s)
	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			occorrenza[c-'a']++
		}
	}

	return occorrenza

}

func stampaOccorrenze(occorrenze [26]int) {
	for i := 0; i < 26; i++ {
		fmt.Printf("%c%3d ", 'a'+i, occorrenze[i]) // %3d crea 3 spazi %c produce il corrispondente simbolo unicode di c
	}
	fmt.Println()
}

func main() {
	s := os.Args                  //s diventa una slice di stringhe
	for i := 1; i < len(s); i++ { //s[i] è una stringa, len(s) è la lunghezza della slice: quanti argomenti possiede
		stampaOccorrenze(contaOccorrenze(s[i]))
	}

}
