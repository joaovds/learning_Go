// Utilizando Structs anônimo

package main

import "fmt"

func main() {
	x := struct {
		nome     string
		esportes []string
		familia  map[string]string
	}{
		nome: "José",
		esportes: []string{
			"basquete",
			"futebol",
			"vôlei",
		},
		familia: map[string]string{
			"irmã": "Julia",
			"Tio":  "Jânio",
			"Tia":  "Jurema",
		},
	}

	fmt.Println(x)
}
