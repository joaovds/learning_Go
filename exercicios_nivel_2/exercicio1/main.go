// Escrever um programa que mostre um número em binário, decimal, e hexadecimal

package main

import "fmt"

func main() {
	x := 27

	fmt.Println("Binário\t\tDecimal\t\tHexadecimal")
	fmt.Printf("%b\t\t%d\t\t%X\t\t", x, x, x)
}
