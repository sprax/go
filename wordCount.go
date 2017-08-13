// https://tour.golang.org/moretypes/23

package main

import (
	"strings"
	"wc" // use local version of wc in src/lib
	// "golang.org/x/tour/wc"		// use external version from GOPATH (~/.golang/lib)
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	for _, word := range strings.Fields(s) {
		m[word] = m[word] + 1
	}
	return m
}

func main() {
	wc.Test(WordCount)
}

func localTest(wc func(s string) map[string]int) {
	m := wc("I ate a donut. Then I ate another donut.")
	print(m)
}
