package main

import (
	"fmt"
)

func bool2str(val bool) string {
	if val {
		return "1"
	} else {
		return "0"
	}
}

func int2str(val int64) string {
	return fmt.Sprintf("%d", val)
}

func float2str(val float64) string {
	return fmt.Sprintf("%.6f", val)
}

func main() {
	fmt.Println("Hello, playground", float2str(3.14159), "what's up", bool2str(false))
    cmd := fmt.Sprintf("add ticket for <@%s|PseudoMike> fix it", "U2YP244MN")
    fmt.Printf("Cmd: <%s>\n", cmd)
}
