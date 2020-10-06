// Criando slice com slices de string (trabalhando com dimensões de slices)

package main

import "fmt"

func main() {
	umaSlice := [][]string{
		[]string{
			"John",
			"Lennon",
			"Música",
		},
		[]string{
			"John",
			"Victor",
			"Programação",
		},
		[]string{
			"John",
			"Teste",
			"Música e programação",
		},
	}

	for _, value := range umaSlice {
		for _, valueSlice2 := range value {
			fmt.Println(valueSlice2)
		}
		fmt.Printf("\n")
	}
}
