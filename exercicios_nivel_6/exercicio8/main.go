// Trabalhando com funções, Métodos, Interfaces

package main

import "fmt"

func main() {
	retornaUmaFuncao()()
}

func retornaUmaFuncao() func() {
	return func() {
		fmt.Println("OIIEEE")
	}
}

// - Crie uma função que retorna uma função.
// - Atribua a função retornada a uma variável.
// - Chame a função retornada.
