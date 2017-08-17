package main

import (
	"fmt"
)

type Session struct {
	playerId string
	beehive  string
	smap     map[string]interface{}
}

func main() {

	session := Session{playerId: "duh", beehive: "bee", smap: make(map[string]interface{})}
	session.smap["horse"] = "Foolish Pride"
	animal := session.smap["duck"]
	if animal == nil {
		animal = session.smap["horse"]
	}
	fmt.Println("Animal:", animal)
}
