/*
	- Utilizando a solução do exercício anterior:
		1. Em package-level scope, utilizando a palavra-chave var, crie uma variável com o identificador "y".
			O tipo desta variável deve ser o tipo subjacente do tipo que você criou no exercício anterior.
    2. Na função main:
*/

package main

import "fmt"

type myType int

var x myType
var y int

func main() {
	fmt.Println(x)        // 1. Demonstre o valor da variável "x"
	fmt.Printf("%T\n", x) // 2. Demonstre o tipo da variável "x"

	x = 42         // 3. Atribua 42 à variável "x" utilizando o operador "="
	fmt.Println(x) // 4. Demonstre o valor da variável "x"

	y = int(x)            // 1. Utilize conversão para transformar o tipo do valor da variável "x" em seu tipo subjacente e, utilizando o operador "=", atribua o valor de "x" a "y"
	fmt.Println(y)        // 2. Demonstre o valor de "y"
	fmt.Printf("%T\n", y) // 3. Demonstre o tipo de "y"
}
