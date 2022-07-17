package main

import "fmt"

var nome string
var nota1, nota2, nota3 int16
var calc_media float32

func main() {
	fmt.Print("Digite o nome do Aluno: ")
	fmt.Scanln(&nome)
	fmt.Print("Digite a primeira nota: ")
	fmt.Scanln(&nota1)
	fmt.Print("Digite a segunda nota: ")
	fmt.Scanln(&nota2)
	fmt.Print("Digite a terceira nota: ")
	fmt.Scanln(&nota3)
	calc_media := (nota1 + nota2 + nota3) / 3
	fmt.Printf("O nome do Aluno: %s\n", nome)
	fmt.Printf("A media do Aluno: %d", calc_media)
}
