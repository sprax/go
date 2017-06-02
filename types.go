
//  http://stackoverflow.com/questions/20170275/how-to-find-a-type-of-a-object-in-golang
    
    // I found 3 ways to recognize type at runtime:

    // Using string formatting
    func typeof(v interface{}) string {
            return fmt.Sprintf("%T", v)
    }

    // Using reflect package
    func typeof(v interface{}) string {
            return reflect.TypeOf(v).String()
    }

package main

import (
    "fmt"
    "reflect"
)

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
    
    fmt.Printf("type assertion N.(type): %v", N.(type))
    
}


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
