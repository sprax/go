package main

import (
	"fmt"
)

func main() {
	var (
		str string
	)
	str = "Hello, playground"
	fmt.Println(str[4:11])
	wds := str.split()

}
