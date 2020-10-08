// Trabalhando com funções, Métodos, Interfaces

package main

import "fmt"

func main() {
	recebeUmaFuncao(funcaoDeArgumento)
}

func funcaoDeArgumento() {
	fmt.Println("Opaaa")
}

func recebeUmaFuncao(x func()) {
	x()
	fmt.Println("Opa de novoo")
}

// - Callback: passe uma função como argumento a outra função.
