package main

import (
    "fmt"
    "strings"
    "math/rand"
    "time"
)

func typoglycemia(s string) string {
    list := strings.Split(s, " ")
    returns := []string{}
    for _, word := range list {
        chars := []rune(word) 
        if len(chars) > 4 {
            last := len(chars) - 1
            word = string(chars[0]) + string(shuffle(chars[1:last])) + string(chars[last])
        }
        returns = append(returns, word)
    }
    return strings.Join(returns, " ")
}

func shuffle(chars []rune) []rune {
    rand.Seed(time.Now().UnixNano())
    n := len(chars)
    for i := n - 1; i >= 0; i-- {
        j := rand.Intn(i + 1)
        chars[i], chars[j] = chars[j], chars[i]
    }
    return chars
}

func main() {
    s := "I couldn't believe that I could actually understand what I was reading : the phenomenal power of the human mind ."
    fmt.Println(typoglycemia(s))
}
