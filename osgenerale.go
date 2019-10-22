package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	s := os.Args[1]

	f, err := os.Open(s)

	if err != nil {
		fmt.Println("erorre", err)
	}

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

}
