// nils.go -- different typed nils in Golang
package main
import "fmt"
func main() {
    fmt.Printf("Func type nil:%#v\n",(func())(nil))
    fmt.Printf("Map type nil:%#v\n", map[string]string(nil))
    fmt.Printf("Slice type nil:%#v\n", []string(nil))
    fmt.Printf("Interface{} type nil:%#v\n", nil)
    fmt.Printf("Channel type nil:%#v\n", (chan struct{})(nil))
    fmt.Printf("Pointer type nil:%#v\n", (*struct{})(nil))
    fmt.Printf("Pointer type nil:%#v\n", (*int)(nil))
    serr := (*string)(nil)
    fmt.Printf("Pointer type nil:%#v\n", serr)
}
// Outputs:
// Func type Nil:(func())(nil)
// Map type Nil:map[string]string(nil)
// Slice type Nil:[]string(nil)
// Interface{} type nil:<nil>
// Channel type nil:(chan struct {})(nil)
// Pointer type nil:(*struct {})(nil)
// Pointer type nil:(*int)(nil)     // same as below
// Pointer type nil:(*string)(nil)  // same as above
