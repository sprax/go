// Time how long it takes to sort an array of N random numbers in Go.
// Usage: go run timeSortRand.go

package main

import "fmt"
import "math/rand"
import "sort"
import "time"

func timeTrack(start time.Time, name string) {
    elapsed := time.Since(start)
    fmt.Printf("%s took %s", name, elapsed)
}



// The prime sieve: Daisy-chain Filter processes.
func main() {
	const NS = 10000000	// number of ints to sort
	const NP = 10		// number of ints to print
	var array [NS]int
	fmt.Println("new: ", array[0:NP])

	//seed := rand.NewSource(time.Now().UnixNano())
	seed := rand.NewSource(0)
	rngA := rand.New(seed)
	for j := 0; j < NS; j++ {
		array[j] = rngA.Intn(NS)
	}
	fmt.Println("rand:", array[0:NP])
	beg := time.Now()
	sort.Ints(array[:])
	elapsed := time.Since(beg)
	fmt.Println("sort:", array[0:NP])
	fmt.Printf("Sorting %d random ints took %s\n", NS, elapsed)
}
