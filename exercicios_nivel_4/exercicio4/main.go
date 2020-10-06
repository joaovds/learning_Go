// Adicionando valores a uma slice com o append

package main

import "fmt"

func main() {
	umaSlice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	umaSlice = append(umaSlice, 52)
	umaSlice = append(umaSlice, 53, 54, 55)

	fmt.Println(umaSlice)

	// Anexando uma slice a outra slice

	outraSlice := []int{56, 57, 58, 59, 60}

	umaSlice = append(umaSlice, outraSlice...)

	fmt.Println(umaSlice)
}
