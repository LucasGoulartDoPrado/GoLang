package main

import "fmt"

var primeiro_valor, segundo_valor float32

func main() {
	fmt.Print("Digite o primeiro valor: ")
	fmt.Scan(&primeiro_valor)
	fmt.Print("Digite o segundo valor: ")
	fmt.Scan(&segundo_valor)
	if primeiro_valor > segundo_valor {
		fmt.Println("O Primeiro valor é maior!")
	} else {
		fmt.Println("O Segundo Valor é maior!")
	}

}
