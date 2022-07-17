package main

import "fmt"

var r float64

func main() {
	pi := 3.14159
	fmt.Print("Informe o valor do Raio: ")
	fmt.Scanln(&r)
	circunferencia := (pi * r) * r
	fmt.Println(circunferencia)
}
