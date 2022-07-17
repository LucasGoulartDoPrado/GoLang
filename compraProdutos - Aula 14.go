package main

import "fmt"

var nomeProdutos, compraProdutos [5]string
var precoProdutos [5]float32
var totalCompra float32
var cont int

func main() {
	for cont = 0; cont < 5; cont++ {
		fmt.Printf("Escreva o nome do %d º produto:", cont+1)
		fmt.Scanln(&nomeProdutos[cont])
		fmt.Printf("Qual o preco do %d º produto: ", cont+1)
		fmt.Scanln(&precoProdutos[cont])
	}
	fmt.Printf("\nProdutos a venda em nossa loja!\n")
	for cont = 0; cont < 5; cont++ {
		fmt.Printf("Produto: %s - R$ %.2f\n", nomeProdutos[cont], precoProdutos[cont])
	}
	fmt.Print("\nEscolha os nosso produtos: ")
	for cont = 0; cont < 5; cont++ {
		fmt.Print("\nEscolha os nosso produtos: ")
		fmt.Scanln(&compraProdutos[cont])
		if compraProdutos[cont] == nomeProdutos[cont] {
			totalCompra = totalCompra + precoProdutos[cont]
		}
	}
	fmt.Printf("Total da Compra: %.2f", totalCompra)
}
