// Removendo uma entrada do Map (excluindo um conteúdo)

package main

import "fmt"

func main() {
	umMap := map[string][]string{
		"Um nome aqui": []string{
			"Um Hobbye aqui", "Outro Hobbye aqui", "E outro Hobbye aqui",
		},
		"Um outro nome aqui": []string{
			"Um Hobbye aqui", "Outro Hobbye aqui", "E outro Hobbye aqui",
		},
	}

	delete(umMap, "johnjohn")

	for indice1, value1 := range umMap {
		fmt.Println(indice1)
		for indice2, value2 := range value1 {
			fmt.Println("\t", indice2, " - ", value2)
		}
	}
}
