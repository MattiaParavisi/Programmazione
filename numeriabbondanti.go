package main

import "fmt"

func main(){
	var quantita int
	fmt.Scan(&quantita)
	for numero:=2; quantita>0 ;numero++{
		//verifico se i è abbondate, se sì, lo stampo
		sommadivisori:=1 //tutti i numeri sono divisibili per 1
		for divisore:=2; divisore<=numero/2 /*mi fermo a i mezzi che è l'ultimo divisore possibile*/; divisore++{ 
			if numero%divisore==0{
				sommadivisori += divisore
			}
		}
		if sommadivisori>numero{
			fmt.Println(numero)
			quantita--
		}
	}
}