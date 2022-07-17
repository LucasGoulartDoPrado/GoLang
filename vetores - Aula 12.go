package main

import "fmt"

var vetSlice = []int{1, 2, 3, 4, 5}
var total, cont, totalPar, totalImpar int

func main() {
	for cont := 0; cont < 5; cont++ {
		fmt.Printf("Digite o %d º numero: ", cont+1)
		fmt.Scan(&vetSlice[cont])
		if (vetSlice[cont] % 2) == 0 {
			totalPar = totalPar + vetSlice[cont]
		} else {
			totalImpar = totalImpar + vetSlice[cont]
		}
		total = total + vetSlice[cont]
	}
	fmt.Printf("Total Par: %d\nTotal Impar: %d\n Total da Soma: %d\n", totalPar, totalImpar, total)
}
