// Recortando uma slice

package main

import "fmt"

func main() {
	umaSlice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	umaSlice = append(umaSlice[:3], umaSlice[len(umaSlice)-4:]...)

	fmt.Println(umaSlice)
}
