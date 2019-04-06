package main

import (
    "fmt"
    "./cabocha"
)

func main() {
    sentenses := cabocha.Load("../neko.txt.cabocha")
    for _, morph := range sentenses[2] {
        fmt.Println(morph)
    }
}
