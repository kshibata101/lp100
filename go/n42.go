package main

import (
    "./cabocha"
    "fmt"
)

func main() {
    sentenses := cabocha.LoadAsChunk("../neko.txt.cabocha")
    for _, sentense := range sentenses {
        for _, chunk := range sentense {
            fmt.Print(cabocha.GetChunkSurfaceWithoutSymbol(chunk))
            if chunk.Dst >= 0 {
                nextChunk := sentense[chunk.Dst]
                fmt.Print("\t")
                fmt.Print(cabocha.GetChunkSurfaceWithoutSymbol(nextChunk))
            }
            fmt.Println()
        }
    }
}
