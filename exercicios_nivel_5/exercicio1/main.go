// Utilizando Structs

package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	sabores   []string
}

func main() {
	pessoa1 := pessoa{
		nome:      "John",
		sobrenome: "Lennon",
		sabores: []string{
			"chocolate",
			"morango",
			"creme",
		},
	}

	pessoa2 := pessoa{
		nome:      "John2",
		sobrenome: "Lennon2",
		sabores: []string{
			"chocolate2",
			"morango2",
			"creme2",
		},
	}

	fmt.Println(pessoa1.nome, pessoa1.sobrenome)
	for _, value := range pessoa1.sabores {
		fmt.Println(value)
	}

	fmt.Println("") // Só para dar uma linha em branco

	fmt.Println(pessoa2.nome, pessoa2.sobrenome)
	for _, value := range pessoa2.sabores {
		fmt.Println(value)
	}
}
