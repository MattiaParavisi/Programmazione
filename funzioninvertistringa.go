package main

import "fmt"

/*func reverse(s string) string{
	var v string
	for i:=len(s)-1; i>=0; i--{
		v+=string(s[i])							//Se scrivo s[i] spara fuori i valori di ogni byte in decimale
	}
	return v
}
*/

func main(){
	var stringa string
	fmt.Scan(&stringa)
	c:=ultrareverse(stringa)
	fmt.Println(c)
}

func ultrareverse(s string) string{
	var v string
	for _, c:=range s{
		v = string(c) + v
		fmt.Println(v)

	}
	return v
}

