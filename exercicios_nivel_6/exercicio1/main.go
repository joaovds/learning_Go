// Trabalhando com funções

package main

import "fmt"

func main() {
	a := retornaInt()
	fmt.Println(a)

	b, c := retornaIntEString()
	fmt.Println(b, c)
}

func retornaInt() int {
	return 27
}

func retornaIntEString() (int, string) {
	return 27, "oii"
}
