package main

import "fmt"

var numero, cont int

func main() {
	for numero >= 0 {
		fmt.Print("Digite o numero: ")
		fmt.Scanln(&numero)
		cont++
	}
	fmt.Println("A quantidade de numeros positivos digitados é: ", cont-1)
}
