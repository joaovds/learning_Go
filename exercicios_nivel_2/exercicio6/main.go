// utilizando iota, crie 4 constantes cujos valores sejam os próximos 4 anos

package main

import "fmt"

const (
	_ = 2020 + iota
	b
	c
	d
	e
)

func main() {
	fmt.Println(b, c, d, e)
}
