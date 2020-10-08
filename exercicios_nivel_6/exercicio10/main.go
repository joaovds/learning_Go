// Trabalhando com funções, Métodos, Interfaces

package main

import "fmt"

func main() {
	x := umaFuncao()
	y := umaFuncao()

	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(y())
	fmt.Println(y())
}

func umaFuncao() func() int {
	a := 0
	return func() int {
		a++
		return a
	}
}

// - Demonstre o funcionamento de um closure.
// - Ou seja: crie uma função que retorna outra função, onde esta outra função faz uso de uma variável alem de seu scope.
