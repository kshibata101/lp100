package main

import (
    "./morpheme"
    "fmt"
)

func main() {
    states := morpheme.Load("../neko.txt.mecab")
    for _, state := range states {
        for i, morph := range state {
            if morph.Pos == "助詞" && morph.Pos1 == "連体化" && i-1 >= 0 && i+1 < len(state) {
                fmt.Println(state[i-1].Surface + state[i].Surface + state[i+1].Surface)
            }
        }
    }
}
