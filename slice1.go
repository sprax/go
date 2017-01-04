// https://tour.golang.org/moretypes/7
package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[2:3]
	fmt.Println(s)
}

