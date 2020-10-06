// Põe na tela o unicode de todas letras maiúsculas 3 vezes cada
// Unicode - Letras maiúsculas: 65 até 90

package main

import "fmt"

func main() {
	for i := 65; i <= 90; i++ {
		fmt.Printf("%c, ", i)
	}
}
