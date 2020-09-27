/*
	- Crie um tipo. O tipo subjacente deve ser int.
	- Crie uma variável para este tipo, com o identificador "x", utilizando a palavra-chave var.
	- Na função main:
*/

package main

import "fmt"

type myType int

var x myType

func main() {
	fmt.Println(x)        // 1. Demonstre o valor da variável "x"
	fmt.Printf("%T\n", x) // 2. Demonstre o tipo da variável "x"

	x = 42         // 3. Atribua 42 à variável "x" utilizando o operador "="
	fmt.Println(x) // 4. Demonstre o valor da variável "x"
}
