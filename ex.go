package main

import "fmt"
//import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	const size = 10
	pic := make([][]uint8, size)
	for j := 0; j < size; j++ {
		pic[j] = make([]uint8, size)
		for k := 0; k < size; k++ {
			pic[j][k] = uint8((j + k)/2)
		}
	}
	return pic
}

func main() {
	pic := Pic(2, 3)
	for _, value := range pic {
		fmt.Printf("%d\n", value)
	}
}

