package main

import (
	"fmt"
)

func sliceOccorrenze(s string) []occorrenza {
	var o []occorrenza
	for _, c := range s {
		found := false

		for i := 0; i < len(o); i++ {
			if o[i].simbolo == c {
				o[i].quanti++
				found = true
				break
			}
		}

		if !found {
			o = append(o, occorrenza{simbolo: c, quanti: 1})
		}

	}
	return o
}

func sString(o occorrenza) string {
	return fmt.Sprintf("%c%3d", o.simbolo, o.quanti) //produce una stringa che contiene il simbolo seguito da un valore decimale che corrisponde al campo quanti dell'occorrenza O
}

func stampaOccorrenze(o []occorrenza) {
	for i := 0; i < len(o); i++ {
		fmt.Printf("%s", sString(o[i]))
	}
}

type occorrenza struct {
	simbolo rune
	quanti  int
}

func main() {
	var s string
	fmt.Scan(&s)
	stampaOccorrenze(sliceOccorrenze(s))

}
