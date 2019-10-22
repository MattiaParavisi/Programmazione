// Stampa una carta da poker
package main

import "fmt"

import "strconv"

func carta(x uint) (string, string, bool) {
	if x == 0 || x > 52 {
		return "", "", false
	}
	var v, t string
	switch { //seme
	case x < 14:
		t = "cuori"
	case x < 27:
		t = "quadri"
	case x < 40:
		t = "fiori"
	default:
		t = "picche"
	}
	switch r := x % 13; r { // valore
	case 0:
		v = "K"
	case 1:
		v = "A"
	case 11:
		v = "J"
	case 12:
		v = "Q"
	default:
		//v = fmt.Sprintf("%v", r)  // r è un int, noi però vogliamo passare alla sua corrispondente forma testuale, per passare da int a string come ItoA
		v=strconv.Itoa(int(r))
	}
	return v, t, true
}

func valoreBJ(val string) (valnum int){
	switch val{
	case "A":
		val=1
	case "J","Q","K":
		val=10
	default:
		valnum, _=strconv.Atoi(val) //ignoro il valore di controllo restituito da Atoi
	}
	return //non scrivo null, viene restituito il valore che è presente in valnum
}

func main() {
	var cartah uint
	fmt.Scan(&cartah)
	j, k, w := carta(cartah)
	
	fmt.Println("valore blackjack")

	if !w{
		fmt.Println("inculati")
		return
	}
	if w {
		fmt.Println(j," di ",k)
	}

	

}
