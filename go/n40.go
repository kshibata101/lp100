package main

import (
    "fmt"
    "./cabocha"
)

func main() {
    sentenses := cabocha.LoadAsMorph("../neko.txt.cabocha")
    for _, morph := range sentenses[2] {
        fmt.Println(morph)
    }
}
