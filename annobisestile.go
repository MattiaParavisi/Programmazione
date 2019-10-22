package main
import "fmt"
func main() {
	var anno int
	fmt.Scan(&anno)
	if anno%4==0 && anno%100!=0||anno%100==0 && anno%400==0 {
		fmt.Println("l'anno è bisestile")
	}else {
		fmt.Println("l'anno non è bisestile")
	}
}
