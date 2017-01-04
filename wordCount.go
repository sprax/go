// https://tour.golang.org/moretypes/23

package main

import (
	"strings"
	"golang.org/x/tour/wc"
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


