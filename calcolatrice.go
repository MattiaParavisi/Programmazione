// Implementa una calcolatrice che esegue ripetutamente operazioni la cui sintassi Ã¨:
// <operatore> <operando>
// dove:
// <operatore> Ã¨ uno fra +,-,*,\, ed S (s),
// <operando> Ã¨ un valore float.
// La calcolatrice visualizza e salva in un registro il risultato di ogni operazione aritmetica.
// Il valore corrente del registro Ã¨ il primo operando (implicito) di una operazione.
// L'operazione
// S x
// assegna al registro il nuovo valore x.
// Se viene digitato un operatore diverso viene stampato il messaggio "operatore sconosciuto".
// Se il divisore Ã¨ 0.0 viene stampato il messaggio "divisione per zero" (in entrambi i casi il programma procede).
// Il comando Q (q) termina il programma.

package main

import "fmt"

func main() {
	var operando,register float64
	operatore:= ""
	for {
		fmt.Scanln(&operatore,&operando)
		switch operatore {
		case "+":
			register+=operando
		case "-":
			register-=operando
		case "/":
			if operando==0{
				fmt.Println("divisione per 0")
				continue
			}
			register/=operando
		case "*":
			register*=operando
		case "S","s":
			register=0
		case "Q","q":
			return
		default:
			fmt.Println("operatore non riconosciuto")
			continue
		}
		fmt.Println(register)
	}
}
