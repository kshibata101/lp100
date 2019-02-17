package main

import (
    "fmt"
    "strings"
)

func ngram(n int, list []string) []string {
    ngram := []string{}
    for i := 0; i < len(list) - n + 1; i++ {
        sum := ""
        for j := 0; j < n; j++ {
            sum += list[i+j]
        }
        ngram = append(ngram, sum)
    }
    return ngram
}

func main() {
    words := strings.Split("I am an NLPer", " ")
    fmt.Println(ngram(2, words))
    chars := []string{"I", "a", "m", "a", "n", "N", "L", "P", "e", "r"}
    fmt.Println(ngram(2, chars))
}
