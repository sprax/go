package main

import (
	"math/rand"
	"time"
)

// Paul's solution provides a simple, general solution.
//
// The question asks for the "the fastest and simplest way". Let's address this. We'll arrive at our final, fastest code in an iterative manner. Benchmarking each iteration can be found at the end of the answer.
//
// All the solutions and the benchmarking code can be found on the Go Playground. The code on the Playground is a test file, not an executable. You have to save it into a file named XX_test.go and run it with go test -bench ..
//
// I. Improvements
//
// 1. Genesis (Runes)
//
// As a reminder, the original, general solution we're improving is this:

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// 2. Bytes
//
// If the characters to choose from and assemble the random string contains only the uppercase and lowercase letters of the English alphabet, we can work with bytes only because the English alphabet letters map to bytes 1-to-1 in the UTF-8 encoding (which is how Go stores strings).
//
// So instead of:
//
// var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
// we can use:
//
// var letters = []bytes("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
// Or even better:
//
// const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
// Now this is already a big improvement: we could achieve it to be a const (there are string constants but there are no slice constants). As an extra gain, the expression len(letters) will also be a const! (The expression len(s) is constant if s is a string constant.)
//
// And at what cost? Nothing at all. strings can be indexed which indexes its bytes, perfect, exactly what we want.
//
// Our next destination looks like this:

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

// 3. Remainder
//
// Previous solutions get a random number to designate a random letter by calling rand.Intn() which delegates to Rand.Intn() which delegates to Rand.Int31n().
//
// This is much slower compared to rand.Int63() which produces a random number with 63 random bits.
//
// So we could simply call rand.Int63() and use the remainder after dividing by len(letterBytes):

func RandStringBytesRmndr(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}

// This works and is significantly faster, the disadvantage is that the probability of all the letters will not be exactly the same (assuming rand.Int63() produces all 63-bit numbers with equal probability). Although the distortion is extremely small as the number of letters 52 is much-much smaller than 1<<63 - 1, so in practice this is perfectly fine.
//
// To make this understand easier: let's say you want a random number in the range of 0..5. Using 3 random bits, this would produce the numbers 0..1 with double probability than from the range 2..5. Using 5 random bits, numbers in range 0..1 would occur with 6/32 probability and numbers in range 2..5 with 5/32 probability which is now closer to the desired. Increasing the number of bits makes this less significant, when reaching 63 bits, it is negligible.
//
// 4. Masking
//
// Building on the previous solution, we can maintain the equal distribution of letters by using only as many of the lowest bits of the random number as many is required to represent the number of letters. So for example if we have 52 letters, it requires 6 bits to represent it: 52 = 110100b. So we will only use the lowest 6 bits of the number returned by rand.Int63(). And to maintain equal distribution of letters, we only "accept" the number if it falls in the range 0..len(letterBytes)-1. If the lowest bits are greater, we discard it and query a new random number.
//
// Note that the chance of the lowest bits to be greater than or equal to len(letterBytes) is less than 0.5 in general (0.25 on average), which means that even if this would be the case, repeating this "rare" case decreases the chance of not finding a good number. After n repetition, the chance that we sill don't have a good index is much less than pow(0.5, n), and this is just an upper estimation. In case of 52 letters the chance that the 6 lowest bits are not good is only (64-52)/64 = 0.19; which means for example that chances to not have a good number after 10 repetition is 1e-8.
//
// So here is the solution:

const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
)

func RandStringBytesMask(n int) string {
	b := make([]byte, n)
	for i := 0; i < n; {
		if idx := int(rand.Int63() & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i++
		}
	}
	return string(b)
}
