// Writing files in Go follows similar patterns to the
// ones we saw earlier for reading.

package main

import (
    "bufio"
    "fmt"
    "io/ioutil"
    "os"
    "reflect"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

    // To start, here's how to dump a string (or just
    // bytes) into a file.
    fn := "/tmp/dat1"
    d1 := []byte("hello\ngo\n")
    err := ioutil.WriteFile(fn, d1, 0644)
    check(err)
    fmt.Printf("ioutil.WriteFile(%s, %v, 0644) dumped file %s\n", fn, d1, fn)

    // For more granular writes, open a file for writing.
    f, err := os.Create("/tmp/dat2")
    check(err)

    fType := reflect.TypeOf(f)
    // It's idiomatic to defer a `Close` immediately
    // after opening a file.
    defer f.Close()

    // You can `Write` byte slices as you'd expect.
    d2 := []byte{115, 111, 109, 101, 32, 98, 121, 116, 101, 115}
    n2, err := f.Write(d2)
    check(err)
    fmt.Printf("%s.Write(%s) wrote %d bytes to a file named %s\n", fType, string(d2), n2, f.Name())

    // A `WriteString` is also available.
    n3, err := f.WriteString("\n" + "this is a string" + "\n")
    fmt.Printf("File.WriteString wrote %d bytes to a file named %s\n", n3, f.Name())

    // Issue a `Sync` to flush writes to stable storage.
    f.Sync()

    // `bufio` provides buffered writers in addition
    // to the buffered readers we saw earlier.
    w := bufio.NewWriter(f)
    wType := reflect.TypeOf(w)
    str := "this is also a string"
    n4, err := w.WriteString(str)
    fmt.Printf("%s.WriteString(%s) wrote %d bytes to %s\n", wType, str, n4, f.Name())

    _, err = w.WriteString("\n")
    _, err = w.WriteString("yet another string")
    _, err = w.WriteString("\n")

    // Use `Flush` to ensure all buffered operations have
    // been applied to the underlying writer.
    w.Flush()
}
