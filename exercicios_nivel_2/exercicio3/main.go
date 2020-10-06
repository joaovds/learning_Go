// Crie constantes tipadas e não tipadas, e mostre seus valores

package main

import "fmt"

const a int = 27
const b = 27

func main() {
	fmt.Printf("%v\t\t%T\n", a, a)
	fmt.Printf("%v\t\t%T\n", b, b)
}
