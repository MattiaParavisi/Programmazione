package main

import (
	"fmt"
	"io"
	"strconv"
)

func main() {
	var res int
	var appendere string
	var err error
	var stringa []string

	for {
		_, err = fmt.Scan(&appendere)
		if err != io.EOF {
			stringa = append(stringa, appendere)
		} else {
			break
		}
	}

	for i := 0; i < len(stringa); i++ {

		switch true {

		case stringa[i] == "+":
			op1, _ := strconv.Atoi(stringa[i-1])
			if res == 0 {
				res += op1
			}
			op1, _ = strconv.Atoi(stringa[i+1])
			res += op1
			restring := strconv.Itoa(res)
			stringa = append([]string{restring}, stringa[i+2:]...)

		case stringa[i] == "-":
			op1, _ := strconv.Atoi(stringa[i-1])
			if res == 0 {
				res += op1
			}
			op1, _ = strconv.Atoi(stringa[i+1])
			res -= op1
			restring := strconv.Itoa(res)
			stringa = append([]string{restring}, stringa[i+2:]...)

		}

		i = 0

	}

	fmt.Println(res)
}
