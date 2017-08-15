package main

import (
	"fmt"
	"strings"
)

func main() {
	hello := "Hello again, dear playground"
	fmt.Println(hello)
	fields := strings.Fields(hello)
	fmt.Printf("Fields: %T  %v\n", fields, fields)
	for idx, fld := range fields {
		fmt.Printf("%d. <%s>\n", idx, fld)
	}
}
