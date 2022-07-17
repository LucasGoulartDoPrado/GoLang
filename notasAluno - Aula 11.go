package main

import "fmt"

func main() {
	var cont, qtdNotas int
	var nota, soma, media float32
	fmt.Printf("Digite a quantidade de notas de um aluno: ")
	fmt.Scan(&qtdNotas)
	for cont = 1; cont <= qtdNotas; cont++ {
		fmt.Printf("Digite a %d º nota: ", cont)
		fmt.Scan(&nota)
		soma = soma + nota
	}
	fmt.Printf("A soma das notas: %.2f\n", soma)
	var qtdNotasConv float32 = float32(qtdNotas)
	media = soma / qtdNotasConv
	fmt.Printf("A media das notas é: %.2f\n", media)
}
