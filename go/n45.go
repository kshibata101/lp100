package main

import (
    "./cabocha"
    "fmt"
    "strings"
)

func main() {
    sentences := cabocha.LoadAsChunk("../neko.txt.cabocha")
    for _, sentence := range sentences {
        for _, chunk := range sentence {
            for _, morph := range chunk.Morphs {
                if morph.Pos == "動詞" {
                    kaku := []string{}
                    for _, src := range chunk.Srcs {
                        prevChunk := sentence[src]
                        for _, prevMorph := range prevChunk.Morphs {
                            if prevMorph.Pos == "助詞" {
                                kaku = append(kaku, prevMorph.Base)
                            }
                        }
                    }
                    fmt.Println(morph.Base + "\t" + strings.Join(kaku, " "))
                }
            }
        }
    }
}
