package main

import (
    "flag"
    "fmt"
    "strings"
)

func main() {
    src := readInput()
    count := countWords(src)
    fmt.Println(count)
}

func readInput() (src string) {
    flag.Parse()
    return strings.Join(flag.Args(), " ")
}

func countWords(src string) int {
    var count int
    for _, word := range strings.Split(src, " ") {
        if word != "" {
            count++
        }
    }
    return count
}
