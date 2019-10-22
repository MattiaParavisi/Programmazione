package main
import "fmt"
func main(){
	var numero,i int
	fmt.Scan(&numero)
	j:=numero
	k:=0
	for i=0; i<numero; i++{
		for d:=0; d<i;d++{
			fmt.Print(" ")
		}
		for d:=0; d<j;d++{
			fmt.Print("* ")
		}
		j--
		fmt.Println("")
	}
	j=numero
	for i=1;i<=j-1;i++{
	
		for d:=2;d<=j-i;d++{
		fmt.Print(" ")
		}

		for k!= i-2{
			fmt.Print("* ")
			k++
		}
		fmt.Println("")

	}
}