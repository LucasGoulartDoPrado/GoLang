package main

import (
	"fmt"
)

var nome, sexo string

func main() {
	fmt.Print("Digite o nome: ")
	fmt.Scanln(&nome)
	fmt.Print("Digite seu sexo - (f-feminino / m-masculino): ")
	fmt.Scanln(&sexo)
	switch {
	case sexo == "f" || sexo == "F":
		fmt.Printf("Seu nome é: %s\n Seu sexo: Feminino!", nome)
	case sexo == "m" || sexo == "M":
		fmt.Printf("Seu nome é: %s\n Seu sexo: Masculino!", nome)
	}
}
