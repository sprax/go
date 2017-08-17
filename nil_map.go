package main

import (
	"fmt"
	"strings"
)

func main() {
	hello := "Hello again, dear playground"

	var table map[string]interface{}

	if nil == (map[string]interface{}(nil)) {
		fmt.Println(hello, "        nil == (map[string]interface{}(nil))")
	}
	if nil == table {
		fmt.Println(hello, "        nil == table")
	}
	fields := strings.Fields(hello)
	fmt.Printf("Fields: %T  %v\n", fields, fields)
	for idx, fld := range fields {
		fmt.Printf("%d. <%s>\n", idx, fld)
	}
}
