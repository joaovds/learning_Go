// Trabalhando com funções, Métodos, Interfaces

package main

import (
	"fmt"
	"math"
)

type quadrado struct {
	lado float64
}

func (q quadrado) area() {
	fmt.Println("Área do quadrado:", (q.lado * q.lado))
}

type circulo struct {
	raio float64
}

func (c circulo) area() {
	fmt.Println("Área do círculo:", (2 * math.Pi * c.raio))
}

type info interface {
	area()
}

func calculaArea(i info) {
	i.area()
}

func main() {
	quadrado1 := quadrado{
		lado: 10,
	}
	quadrado1.area()

	circulo1 := circulo{
		raio: 10,
	}
	circulo1.area()

	calculaArea(quadrado1)
	calculaArea(circulo1)
}

// - Crie um tipo "quadrado"
// - Crie um tipo "círculo"
// - Crie um método "área" para cada tipo que calcule e retorne a área da figura
//     - Área do círculo: 2 * π * raio
//     - Área do quadrado: lado * lado
// - Crie um tipo "figura" que defina como interface qualquer tipo que tiver o método "área"
// - Crie uma função "info" que receba um tipo "figura" e retorne a área da figura
// - Crie um valor de tipo "quadrado"
// - Crie um valor de tipo "círculo"
// - Use a função "info" para demonstrar a área do "quadrado"
// - Use a função "info" para demonstrar a área do "círculo"
