package main

import (
    "./morpheme"
    "fmt"
)

func main() {
    filepath := "../neko.txt.mecab"
    states := morpheme.Load(filepath)

    nouns_list := []string{}
    for _, state := range states {
        nouns := ""
        for _, morpheme := range state {
            if morpheme.Pos == "名詞" {
                nouns += morpheme.Surface
            } else {
                if len(nouns) > 0 {
                    nouns_list = append(nouns_list, nouns)
                }
                nouns = ""
            }
        }
        if len(nouns) > 0 {
            nouns_list = append(nouns_list, nouns)
        }
    }
    fmt.Println(nouns_list)
}
