package main

import (
	"fmt"
)

type medicinali struct {
	nomemedicina   string
	quantitassunta int
}

type maxmedicinali struct {
	nomemedicina string
	quantitamax  int
}

type pazienti struct {
	nome       string
	cognome    string
	assunzioni []medicinali
}

func killer(a []pazienti, b []maxmedicinali) {
	var flag bool
	for i := 0; i < len(a); i++ {
		for c:= 0; c< len(a); c++{
			if b[c].quantitamax < a[i].assunzioni[c].quantitassunta {
				flag=true
				if flag{
					fmt.Println(a[i],"morto")
					break
				}
			} 
		
		}
	}

}

func main() {
	var medicinalimass []maxmedicinali
	var pazientitot []pazienti
	pazientitot = []pazienti{{"a", "_", []medicinali{{"azzurro", 39},{"rosa",11},{"verde",0}}}, 
								{"b", "__", []medicinali{{"azzurro", 41},{"rosa",0},{"verde",0}}}, 
									{"c", "___", []medicinali{{"azzurro", 39},{"rosa",0},{"verde",0}}},
								}
								
								
	medicinalimass = []maxmedicinali{{"azzurro", 40},{"rosa",10},{"verde",50},}

	killer(pazientitot, medicinalimass)
}
