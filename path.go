// Stampa le parti di un path, ad es. se il path Ã¨  "/home/lorenzo/lab" stampa
//home
//lorenzo
//lab
// Se il path non contiene alcuna occorrenza di '/' viene stampato cosÃ¬ com'Ã¨.
package main

import (
	"fmt"
	"strings"
)
const sep = "/"
func main() {
	var stringa string
	var from,to int
	fmt.Scan(&stringa)
	for {
		to = strings.Index(stringa[from:],sep)
		if to<0{ 
			//se non c'e' viene restituito un valore minore di 0
			break
		}
		fmt.Println(stringa[from:to+from])
		from +=1+to 
	}
	fmt.Println(stringa[from:])

	
}
