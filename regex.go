
package main

import (
    "fmt"
    "regexp"
)

func main() {
    matched, err := regexp.MatchString("(?i)add\\s+(a\\s+)?ticket\\s+for\\s.*", "Add a ticket for @godzilla ")
    fmt.Println(matched, err)
    matched, err = regexp.MatchString("(?i)respond\\s+to\\s.*", "Respond  to @bigfoot")
    fmt.Println(matched, err)
    matched, err = regexp.MatchString("a(b", "seafood")
    fmt.Println(matched, err)
}


