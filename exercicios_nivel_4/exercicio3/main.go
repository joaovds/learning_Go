// Fatiando slice... 1° - 3°, 5° - último, 2° - 7°, 3° ao penúltimo

package main

import "fmt"

func main() {
	umaSlice := []int{07, 17, 27, 37, 47, 57, 67, 77, 87, 97}

	fmt.Println(umaSlice[:3])
	fmt.Println(umaSlice[4:len(umaSlice)])
	fmt.Println(umaSlice[1:7])
	fmt.Println(umaSlice[2:(len(umaSlice) - 1)])
}
