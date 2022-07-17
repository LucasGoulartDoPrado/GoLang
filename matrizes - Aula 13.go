package main

import "fmt"

var total, contLinha, contColuna int
var mat [2][5]int

func main() {
	for contLinha = 0; contLinha < 2; contLinha++ {
		for contColuna = 0; contColuna < 5; contColuna++ {
			fmt.Printf("Digite o numero para a linha %d e coluna %d :", contLinha+1, contColuna+1)
			fmt.Scan(&mat[contLinha][contColuna])
			total = total + mat[contLinha][contColuna]
		}
	}
	fmt.Printf("Total da Soma: %d", total)
}
