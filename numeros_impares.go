package main

import "fmt"

var cont int

func main() {
	cont := 0
	for cont <= 20 {
		if cont%2 == 1 {
			fmt.Println(cont)
		}
		cont++
	}
}
