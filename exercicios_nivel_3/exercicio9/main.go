// Criar um programa que use "switch" com switch statement declarado numa var string "esporteFavorito"

package main

import "fmt"

func main() {
	esportFavorito := "Basquete"

	switch esportFavorito {
	case "Basquete":
		fmt.Println("Você gosta de Basquete! hee")
	case "Tênis":
		fmt.Println("Você gosta de Tênis!")
	case "Vôlei":
		fmt.Println("Você gosta de Vôlei!")
	default:
		fmt.Println("De que você gosta mermão?")
	}
}
