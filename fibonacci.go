// Go closure for generating Fibonacci numbers 
// Based on The Go Playground: https://play.golang.org/p/AZaI_hrPgV 
// Usage: go run fibonacci.go

package main

import "fmt"

// fibonacci returns a function that returns
// successive Fibonacci numbers.
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	fib := fibonacci()
	for i := 0; i < 48; i++ {
		if i % 8 == 0 {
			fmt.Print("\n")
		}
		fmt.Printf("%10d, ", fib())
	}
}
