// Mostrar os anos desde meu nascimento até atualmente com laço infinito

package main

import "fmt"

func main() {
	anoMeuNascimento := 2004
	anoAtual := 2020

	for {
		if anoMeuNascimento > anoAtual {
			break
		}
		fmt.Println(anoMeuNascimento)
		anoMeuNascimento++
	}
}
