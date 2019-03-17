package main

import (
    "./morpheme"
    "fmt"
)

func main() {
    states := morpheme.Load("../neko.txt.mecab")
    for _, state := range states {
        for _, morph := range state {
            if morph.Pos1 == "サ変接続" {
                fmt.Println(morph)
            }
        }
    }
}
