// https://tour.golang.org/moretypes/7
package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13, 17, 19}

	var s []int = primes[4:8]
	fmt.Println(s)
}

