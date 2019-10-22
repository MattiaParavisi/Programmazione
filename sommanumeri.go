package main
import "fmt"
func main (){
	var numero int
	fmt.Println("inserisci un numero")
	fmt.Scan(&numero)
	primacifra:= numero%10
	secondacifra:= ((numero%100)-primacifra)/10
	terzacifra:=((numero%1000)-secondacifra-primacifra)/100	
	sommacifre:= primacifra+secondacifra+terzacifra
	if sommacifre<=12{
		fmt.Println("la somma delle cifre del tuo numero è minore di 12")
	}else {
		fmt.Println("la somma delle cifre del tuo numero è maggiore di 12")
	} 
}
