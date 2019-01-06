package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	var campo [6][7]int //[6]righe [7]elementi per riga
	var turno bool
	turno = true
	controllo := 0
	riga := 0
	colonna := 0
	fmt.Println(turno)
	for controllo != 10 {
		fmt.Scan(&riga)
		fmt.Scan(&colonna)
		fmt.Scan(&controllo)
		if turno == true {
			if campo[riga][colonna] == 5 || campo[riga][colonna] == 9 {
				fmt.Println("Errore, Casella già occupata")
				continue
			}
			campo[riga][colonna] = 5

		}
		if turno == false {
			if campo[riga][colonna] == 5 || campo[riga][colonna] == 9 {
				fmt.Println("Errore, Casella già occupata")
				continue
			}
			campo[riga][colonna] = 9

		}
		for i := 0; i < 6; i++ {
			for k := 0; k < 7; k++ {
				fmt.Print(campo[i][k])
			}
			fmt.Println()
		}

		if turno == true {
			turno = false
			continue
		}

		if turno == false {
			turno = true
			continue
		}
	}
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
