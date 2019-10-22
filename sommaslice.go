package main

import (
	"fmt"
	"os"
	"strconv"
)

func somma(a []int) {
	for i := 1; i < len(a); i++ {
		a[i] += a[i-1]
	}
}

func main() {
	s := os.Args
	var a []int
	for i := 1; i < len(s); i++ { //quando usi Args utilizza sempre un bel for che parte da uno sennò ti piglia anche il nome del programma
		x, err := strconv.Atoi(s[i])
		if err != nil { //se è nil va tutto alla grandissima, senno c'è un errore
			fmt.Println("erroraccio")
			return
		}
		a = append(a, x)
	}
	fmt.Println("a prima", a)
	somma(a)
	fmt.Println("a dopo", a)
}
