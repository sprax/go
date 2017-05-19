
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

    addTicketRE := regexp.MustCompile(`(?i)add\s+(a\s+)?ticket\s+for\s.*`)
    matched := addTicketRE.MatchString("Add a ticket for @godzilla ")
    fmt.Println(matched)
    matched, err := regexp.MatchString("(?i)respond\\s+to\\s.*", "Respond  to @bigfoot")
    fmt.Println(matched, err)
    matched, err = regexp.MatchString("a(b", "seafood")
    fmt.Println(matched, err)

    re := regexp.MustCompile("a(x*)b(y|z)c")
    fmt.Printf("%q\n", re.FindStringSubmatch("-axxxbyc-"))
    fmt.Printf("%q\n", re.FindStringSubmatch("-abzc-"))
    
    myText := "Add a  TickeT for @bruce Are you the batman?"
    ticketOnBehalfRegex := regexp.MustCompile(`(?i)add\s+(?:a\s+)?ticket\s+for\s+(.*)`)
    matches := ticketOnBehalfRegex.FindStringSubmatch(myText)
    fmt.Printf("ticketOnBehalfRegex ~=? %s:\n\t  %#v\n", myText, matches)
    
    myText = "ResPond to @alfred Where's Bruce?"
    questionOnBehalfRegex := regexp.MustCompile(`(?i)respond\s+to\s+(?-i)((<@[^|]{5,}\|[a-z0-9][a-z0-9._-]*>)|@[a-z0-9][a-z0-9._-]*)\s+(.*)`)
    matches = questionOnBehalfRegex.FindStringSubmatch(myText)
    fmt.Printf("questionOnBehalfRegex ~=? %s:\n\t  %#v\n", myText, matches)

    myText = "ResPond to <@U54321ABC|rosencrantz> Where's Guildenstern?"
    matches = questionOnBehalfRegex.FindStringSubmatch(myText)
    fmt.Printf("questionOnBehalfRegex ~=? %s:\n\t  %#v\n", myText, matches)

    myText = "ResPond to <@U54321ABC|rosencrantz> Where's Guildenstern?"
    matches = questionOnBehalfRegex.FindStringSubmatch(myText)
    fmt.Printf("questionOnBehalfRegex ~=? %s:\n\t  %#v\n", myText, matches)
    
    myText = "ResPond <@U54321ABC|rosencrantz> Where's Rosencrantz?"
    matches = questionOnBehalfRegex.FindStringSubmatch(myText)
    fmt.Printf("questionOnBehalfRegex ~=? %s:\n\t  %#v\n", myText, matches)
    if (matches != nil && ) {
        fmt.Println("matches is true!  ", matches)
    } else {
        fmt.Println("matches is false!  ", matches)
    }
}


