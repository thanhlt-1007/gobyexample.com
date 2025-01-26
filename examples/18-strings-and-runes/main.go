package main

import (
    "fmt"
    "unicode/utf8"
)

func examineRune(r rune) {
    if r == 't' {
        fmt.Println("found tee")
    } else if r == 'ส' {
        fmt.Println("found so sua")
    }
}

func main() {
    const s = "สวัสดี"

    fmt.Println("Len:", len(s))

    for i:= 0; i < len(s); i++ {
        fmt.Printf("%d: %x\n",i, s[i])
    }

    fmt.Println("Rune count: ", utf8.RuneCountInString(s))

    for i, runeValue := range s {
        fmt.Printf("%#U starts at %d\n", runeValue, i)
    }

    fmt.Println("\nUsing DecodeRuneInString")
    for i, w := 0, 0; i < len(s); i += w {
        runeValue, width := utf8.DecodeLastRuneInString(s[i:])
        fmt.Printf("%#U starts at %d\n", runeValue, i)
        w = width
        examineRune(runeValue)
    }
}
