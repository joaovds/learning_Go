// Mostrar os anos desde meu nascimento até atualmente

package main

import "fmt"

func main() {
	anoMeuNascimento := 2004
	anoAtual := 2020

	for anoMeuNascimento <= anoAtual {
		fmt.Println(anoMeuNascimento)
		anoMeuNascimento++
	}
}
