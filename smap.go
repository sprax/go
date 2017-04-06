package main

import (
	"fmt"
	"reflect"
)

type Play struct {
	Smap map[string]interface{}
	Hour int
}

func main() {
	fmt.Println("Hello, playground")

	play := Play {
		Hour : 4,
		Smap : map[string]interface{} {},
	}

	fmt.Printf("play: %#v\n", play)

	tst := "string"
	tst2 := 10
	tst3 := 1.2

	fmt.Println(reflect.TypeOf(tst))
	fmt.Println(reflect.TypeOf(tst2))
	fmt.Println(reflect.TypeOf(tst3))

	fmt.Println("Later, playground")
}
