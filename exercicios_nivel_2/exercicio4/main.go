package main

import "fmt"

func main() {
	a := 27
	fmt.Println("Binário\t\tDecimal\t\tHexadecimal")
	fmt.Printf("%b\t\t%d\t\t%#X\n", a, a, a)

	b := a << 1
	fmt.Printf("%b\t\t%d\t\t%#X\n", b, b, b)
}
