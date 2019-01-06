package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a []string
	var b string
	var accumulatore []int
	var risultato []int
	for {
		fmt.Scan(&b)
		if b == "Stop" {
			break
		}
		a = append(a, b)
	}

	for _, c := range a {
		atoi, err := strconv.Atoi(string(c))
		if err == nil {
			accumulatore = append(accumulatore, atoi)
		}

		if c == "+" {
			ris := 0
			for i := 0; i < len(accumulatore); i++ {
				ris += accumulatore[i]
			}
			risultato = append(risultato, ris)
			fmt.Println(risultato)

		}

		if c == "-" {
			ris := accumulatore[0]
			for i := 1; i < len(accumulatore); i++ {
				ris -= accumulatore[i]
			}
			risultato = append(risultato, ris)
			fmt.Println(risultato)
		}

		if c == "/" {
			ris := accumulatore[0]
			for i := 1; i < len(accumulatore); i++ {
				ris = ris / accumulatore[i]
			}
			risultato = append(risultato, ris)
			fmt.Println(risultato)
		}

		if c == "*" {
			ris := 1
			for i := 0; i < len(accumulatore); i++ {
				ris *= accumulatore[i]
			}
			risultato = append(risultato, ris)
			fmt.Println(risultato)
		}

	}

}
