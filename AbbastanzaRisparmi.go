package main

import "fmt"

func Soluzione(risparmi []int) int {
	giorni := 0
	somma := 0

	if len(risparmi) == 0 {
		fmt.Println("Entro in len")
		return 0
	}

	if risparmi[0] == 0 {
		return 0
	}

	for i := 1; i < len(risparmi); i++ {
		somma += risparmi[i]
		if somma < risparmi[0] {
			giorni++
			continue
		} else if somma >= risparmi[0] {
			giorni++
			return giorni
		}
	}

	return 0

}
func main() {
	risparmi := []int{}
	aggiunta := 1
	fmt.Println("Inserisci una sequenza di numeri, terminata da 0, in cui in prima posizione si trova l'oggetto desiderato e poi i risparmi")
	for aggiunta != 0 {
		fmt.Scan(&aggiunta)
		risparmi = append(risparmi, aggiunta)
	}
	sol := Soluzione(risparmi)
	fmt.Println(sol)
}
