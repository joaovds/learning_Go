// Criar um programa que demostre o funcionamento do "if" junto com "else if" and "else"

package main

import "fmt"

func main() {
	idade := 27

	if (idade >= 18) && (idade < 60) {
		fmt.Println("Você é maior de idade")
	} else if idade >= 60 {
		fmt.Println("Você têm experiência!")
	} else {
		fmt.Println("Você é menor de idade")
	}
}
