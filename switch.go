// https://tour.golang.org/flowcontrol/9

package main

import (
	"fmt"
	"runtime"
)


// Go has built-in support for _multiple return values_.
// This feature is used often in idiomatic Go, for example
// to return both result and error values from a function.

// The `(int, int)` in this function signature shows that
// the function returns 2 `int`s.
func vals() (int, int) {
        return 3, 7
}

func dork() (string, error) {
        return "duh", nil
}

func switcher() {

        // Here we use the 2 different return values from the
        // call with _multiple assignment_.
        a, b := vals()
        fmt.Println(a)
        fmt.Println(b)

        // If you only want a subset of the returned values,
        // use the blank identifier `_`.
        _, c := vals()
        fmt.Println(c)

        ss, err := dork();
        if (err != nil) {
            fmt.Println(ss)
        }
}

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
    switcher()
}

