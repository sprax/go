package main

import "fmt"

/**
 * Ruby code below. How to do equivalent in Go?
people = [
  {
    first: "Nic",
    last: "Williams"
  },
  {
    first: "Banjo",
    last: "Williams"
  }
]
p people.map { |person| person["first"] }
*/

type Person struct {
	first string
	last  string
}

type People []Person

func NewPeople() People {
	return People{}
}

// Explicit method
func (people People) firstNames() []string {
	firstNames := make([]string, len(people))

	for i, person := range people {
		fmt.Println(person.first)
		firstNames[i] = person.first
	}
	return firstNames
}

// Reusable enumeration method
func (people People) Map(f func(*Person) string) []string {
	res := make([]string, len(people))
	for i, person := range people {
		res[i] = f(&person)
	}
	return res
}

func main() {
	people := People{
		{"Nic", "Williams"},
		{"Banjo", "Williams"},
	}
	proc := func(p *Person) string {
		return p.first
	}
	fmt.Printf("%q\n", people.Map(proc))

	//fmt.Printf("%#v\n", people)
	//fmt.Println(people.firstNames())
}
