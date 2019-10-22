package main
import "fmt"
func main(){
	var numero int
	fmt.Println("Inserisci un numero")
	fmt.Scan(&numero)
	if numero%1000==0{
		fmt.Println("il numero finisce con 000 ")
	}else {
		fmt.Println("il numero NON finisce con 000 ")
	}
}
