// Utilizando o range e slice

package main

import "fmt"

func main() {
	umaSlice := []int{07, 17, 27, 37, 47, 57, 67, 77, 87, 97}

	for _, value := range umaSlice {
		fmt.Printf("%v\t%T\n", value, value)
	}
}
