package main

import "fmt"

type persona struct {
	nome           string
	cognome        string
	numerotelefono int
}

func cambionome(n []persona) {
	for i := 0; i < len(n); i++ {
		n[i].nome = "carletto"
	}
}

func main() {
	citta := []persona{
		persona{"franco", "cavalli", 3206332334},
		persona{"giulio", "cesare", 3458223111},
		persona{"carlo", "magno", 786576474747},
	}

	fmt.Println(citta)

	cambionome(citta)

	fmt.Println(citta)

}
