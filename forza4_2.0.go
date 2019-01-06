package main

import (
	"fmt"
	"os"
	"os/exec"
)

func vittoria(tabellone [6][7]int) (int, bool) {
	//condizione orizzontale
	for i := 0; i < 6; i++ {
		for k := 0; k < 3; k++ {
			if tabellone[i][k] == 5 && tabellone[i][k+1] == 5 && tabellone[i][k+2] == 5 && tabellone[i][k+3] == 5 {
				return 0, true
			} else if tabellone[i][k] == 9 && tabellone[i][k+1] == 9 && tabellone[i][k+2] == 9 && tabellone[i][k+3] == 9 {
				return 1, true
			}
		}
	}

	//condizione verticale
	for k := 0; k < 7; k++ {
		for i := 5; i > 3; i-- {
			if tabellone[i][k] == 5 && tabellone[i-1][k] == 5 && tabellone[i-2][k] == 5 && tabellone[i-3][k] == 5 {
				return 0, true
			} else if tabellone[i][k] == 9 && tabellone[i-1][k] == 9 && tabellone[i-2][k] == 9 && tabellone[i-3][k] == 9 {
				return 1, true
			}
		}
	}
	//condizione diagonale
	for i := 5; i > 3; i-- {
		for k := 0; k < 3; k++ {
			if tabellone[i][k] == 5 && tabellone[i-1][k+1] == 5 && tabellone[i-2][k+2] == 5 && tabellone[i-3][k+3] == 5 {
				return 0, true
			} else if tabellone[i][k] == 9 && tabellone[i-1][k+1] == 9 && tabellone[i-2][k+2] == 9 && tabellone[i-3][k+3] == 9 {
				return 1, true
			}
		}
	}

	return 2, false

}

func printtabellone(tabellone [6][7]int) {
	for i := 0; i < 6; i++ {
		for k := 0; k < 7; k++ {
			fmt.Print(tabellone[i][k])
		}
		fmt.Println()
	}
}

func creatabellone() [6][7]int { //[6]righe [7]colonna verticale
	var tabellone [6][7]int
	return tabellone
}

func checkcolonna(colonna int) bool {
	if colonna < 7 && colonna >= 0 {
		return true
	}
	return false
}

func checkoccupata(tabellone [6][7]int, colonna int) (int, bool) {
	var riga int
	for riga = 5; riga >= 0; riga-- {
		if tabellone[riga][colonna] == 0 || tabellone[riga][colonna] == 0 {
			return riga, true
		}
	}
	return 0, false
}

func addpedina(turno int, tabellone [6][7]int, colonna int) [6][7]int {
	errigacolonna := checkcolonna(colonna)
	if turno > 1 {
		clearscreen()
	}
	riga, erroccupata := checkoccupata(tabellone, colonna)
	if errigacolonna == true && erroccupata == true {
		if turno%2 != 0 {
			tabellone[riga][colonna] = 5
		} else if turno%2 == 0 {
			tabellone[riga][colonna] = 9
		}
		return tabellone
	}

	return tabellone
}

func clearscreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	var colonna, giocvitt int
	vitt := false
	turno := 1
	tabellone := creatabellone()
	for vitt != true {
		fmt.Scan(&colonna)
		tabellone = addpedina(turno, tabellone, colonna)
		printtabellone(tabellone)
		giocvitt, vitt = vittoria(tabellone)
		turno++
	}

	if giocvitt == 0 {
		fmt.Println("Vince il 5")
	} else if giocvitt == 1 {
		fmt.Println("Vince il 9")
	}

}
