// go run types.go

package main

import (
    "fmt"
    "reflect"
)

// Using type assertions
func typeof(v interface{}) string {
    switch t := v.(type) {
        case int:
            return "int"
        case float64:
            return "float64"
        //... etc
        default:
            _ = t
        return "unknown"
    }
}

func main() {
    fmt.Println("Hello, playground")
    var (
        N int
        S string
    )
    N = 2
    S = "two"
    
    fmt.Printf("N: %T,  S:  %T\n", N, S)
    
    fmt.Printf("reflect.TypeOf(N):  %s\n", reflect.TypeOf(N) )
    
    fmt.Printf("type assertion switch as func typeof: %v\n", typeof(N))
}

