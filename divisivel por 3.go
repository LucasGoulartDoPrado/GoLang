package main

import (
	"fmt"
)

var n int

func main() {
	n := 3
	for n != 0 {
		fmt.Print("Digite o número: ")
		fmt.Scan(&n)
		if n%3 == 0 {
			fmt.Println("É divisivel por 3!")
		} else {
			fmt.Println("Não é divisivel por 3!")
		}
	}
}
