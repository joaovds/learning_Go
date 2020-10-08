// Trabalhando com funções, Métodos

package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func (p pessoa) mostrarPessoa() {
	fmt.Println("Nome Completo:", p.nome, p.sobrenome, "\nIdade:", p.idade)
}

func main() {
	pessoa1 := pessoa{
		nome:      "joão",
		sobrenome: "Victor",
		idade:     16,
	}

	pessoa1.mostrarPessoa()
}

// - Crie um tipo struct "pessoa" que contenha os campos:
//     - nome
//     - sobrenome
//     - idade
// - Crie um método para "pessoa" que demonstre o nome completo e a idade;
// - Crie um valor de tipo "pessoa";
// - Utilize o método criado para demonstrar esse valor.
