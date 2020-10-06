// Usando o make para o slice, e mostrando o tamanho (len) e capacidade (cap) do slice

package main

import "fmt"

func main() {
	estadosBr := make([]string, 26, 26)

	estadosBr = []string{"Acre", "Alagoas", "Amapá", "Amazonas", "Bahia", "Ceará", "Espírito Santo", "Goiás", "Maranhão", "Mato Grosso", "Mato Grosso do Sul", "Minas Gerais", "Pará", "Paraíba", "Paraná", "Pernambuco", "Piauí", "Rio de Janeiro", "Rio Grande do Norte", "Rio Grande do Sul", "Rondônia", "Roraima", "Santa Catarina", "São Paulo", "Sergipe", "Tocantins"}

	fmt.Println(len(estadosBr), cap(estadosBr))

	// Mostrar todos os valores do slice sem utilizar o range
	for i := 0; i < len(estadosBr); i++ {
		fmt.Println(estadosBr[i])
	}
}
