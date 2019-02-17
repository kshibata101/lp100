package main

import (
    "fmt"
    "strings"
)

func split_to_words_length(s string) []int {
    list := strings.Split(s, " ")
    lengths := make([]int, 0, 0)
    for _, word := range list {
        word = strings.Replace(word, ",", "", -1)
        word = strings.Replace(word, ".", "", -1)
        fmt.Println(word)
        lengths = append(lengths, len(word))
    }
    return lengths
}

func main() {
    s := "Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics."
    fmt.Println(split_to_words_length(s))
}
