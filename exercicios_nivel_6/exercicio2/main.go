// Trabalhando com funções

package main

import "fmt"

func main() {
	umaSlice := []int{
		10, 3, 2, 5,
	}
	soma := parametroVariadico(umaSlice...)
	fmt.Println(soma)

	fmt.Println(parametroSlice(umaSlice))
}

func parametroVariadico(x ...int) int {
	soma := 0
	for _, value := range x {
		soma += value
	}
	return soma
}

func parametroSlice(x []int) int {
	soma := 0
	for _, value := range x {
		soma += value
	}
	return soma
}

// - Crie uma função que receba um parâmetro variádico do tipo int e retorne a soma de todos os ints recebidos;
// - Passe um valor do tipo slice of int como argumento para a função;
// - Crie outra função, esta deve receber um valor do tipo slice of int e retornar a soma de todos os elementos da slice;
// - Passe um valor do tipo slice of int como argumento para a função.
