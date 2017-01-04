// From https://tour.golang.org/flowcontrol/8

package main

import (
	"fmt"
)

// Newton's method, fixed iterations
func Sqrt(x float64) float64 {
    z := x/2.0;
    for j := 0; j < 10; j++ {
		z = z - (z*z - x)/(2.0*z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}

