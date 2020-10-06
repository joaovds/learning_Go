// Criar um programa que use "switch" sem o switch statement declarado

package main

import "fmt"

func main() {
	idade := 33

	switch {
	case (idade >= 18) && (idade < 60):
		fmt.Println("Você é maior de idade")
	case (idade >= 60):
		fmt.Println("Você têm experiência!")
	default:
		fmt.Println("Você é menor de idade")
	}
}
