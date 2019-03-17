package main

import (
    "./morpheme"
    "fmt"
)

func main() {
    states := morpheme.Load("../neko.txt.mecab")
    fmt.Println(states)
}
