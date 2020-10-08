// Utilizando Structs

package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	sabores   []string
}

func main() {

	meuMap := make(map[string]pessoa)

	meuMap["John"] = pessoa{
		nome:      "John",
		sobrenome: "Lennon",
		sabores: []string{
			"chocolate",
			"morango",
			"creme",
		},
	}

	meuMap["John2"] = pessoa{
		nome:      "John2",
		sobrenome: "Lennon2",
		sabores: []string{
			"chocolate2",
			"morango2",
			"creme2",
		},
	}

	for _, value := range meuMap {
		fmt.Println(value.nome, value.sobrenome)
		for _, value2 := range value.sabores {
			fmt.Println(value2)
		}
		fmt.Println("")
	}
}
