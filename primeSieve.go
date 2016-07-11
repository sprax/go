// A concurrent prime sieve
// Based on The Go Playground: https://play.golang.org/p/9U22NfrXeq
// Usage: go run primeSieve.go

package main

import "fmt"

// Send the sequence 2, 3, 4, ... to channel 'ch'.
func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}

// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in // Receive value from 'in'.
		if i%prime != 0 {
			out <- i // Send 'i' to 'out'.
		}
	}
}

// The prime sieve: Daisy-chain Filter processes.
func main() {
	ch := make(chan int) // Create a new channel.
	go Generate(ch)      // Launch Generate goroutine.
	k := 0
	for i := 0; i < 9590; i++ {
		prime := <-ch
		fmt.Printf("%6d ", prime)
		ch1 := make(chan int)
		go Filter(ch, ch1, prime)
		ch = ch1
		k++
		if k == 10 {
			k = 0
			fmt.Print("\n")
		}
	}
}

