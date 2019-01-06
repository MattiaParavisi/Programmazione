package main

import "fmt"

type Pezzo struct {
	colore bool
}

//Definisco una casella
type Casella struct {
	riga    int
	colonna int
}

type Tabellone map[Casella]Pezzo

func main() {
	var counttab, riga, r, c, turnoperaltro int //r è la casella //c è la riga
	var p Pezzo
	var s Tabellone
	var turno bool
	s = make(map[Casella]Pezzo)
	turnoperaltro = 0
	turno = true
	lettera := 'a'
	for stampalettera := 0; stampalettera < 7; stampalettera++ {
		fmt.Print("   ", string(lettera))
		lettera++
	}
	fmt.Println()
	fmt.Scan(&r, &c)
	numero := 1
	for counttab = 1; counttab < 7; counttab++ {
		fmt.Print(numero)
		numero++
		for riga = 1; riga <= 7; riga++ {
			fmt.Print("| ")
			if r == riga && c == counttab && turno == true {
				fmt.Print("X ")
				turnoperaltro++
				p.colore = true
				turno = false
				continue
			}
			if r == riga && c == counttab && turno == false {
				fmt.Print("O ")
				turnoperaltro++
				p.colore = false
				turno = true
				continue
			}
			fmt.Print("__")
		}
		fmt.Print("|")
		fmt.Println()
	}

	s[Casella{r, c}] = p

	fmt.Print(s)
}
