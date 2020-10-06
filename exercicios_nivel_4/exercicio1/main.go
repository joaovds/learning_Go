// Utilizando o range em array

package main

import "fmt"

func main() {
	umaArray := [5]int{07, 17, 27, 37, 47}

	for _, value := range umaArray {
		fmt.Printf("%v\t%T\n", value, value)
	}
}
