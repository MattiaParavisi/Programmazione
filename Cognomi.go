package main

import "fmt"

func Soluzionee(runetta string) (int, int) {
	minL := 0
	magM := 0
	j := 0

	if len(runetta) == 0 {
		return 0, 0
	}

	for i := 0; i < len(runetta); i++ {

		if runetta[i] <= byte('L') {
			minL++
			j = i
			for runetta[j] != byte('*') && j != len(runetta)-1 {
				j++
			}
			i = j
		}

		if runetta[i] >= byte('M') {
			magM++
			j = i
			for runetta[j] != byte('*') && j != len(runetta)-1 {
				j++
			}
			i = j
		}
	}

	return minL, magM
}

func main() {
	var runetta string
	fmt.Scan(&runetta)
	sol1, sol2 := Soluzionee(runetta)
	fmt.Println(sol1, sol2)
}
