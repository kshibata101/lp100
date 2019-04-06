package main

import (
    "./cabocha"
    "fmt"
)

func main() {
    sentenses := cabocha.LoadAsChunk("../neko.txt.cabocha")
    for _, sentense := range sentenses {
        for _, chunk := range sentense {
            if chunk.Dst >= 0 && cabocha.HasPos(chunk, "名詞") {
                nextChunk := sentense[chunk.Dst]
                if cabocha.HasPos(nextChunk, "動詞") {
                    fmt.Print(cabocha.GetChunkSurfaceWithoutSymbol(chunk))
                    fmt.Print("\t")
                    fmt.Print(cabocha.GetChunkSurfaceWithoutSymbol(nextChunk))
                    fmt.Println()
                }
            }
        }
    }
}
