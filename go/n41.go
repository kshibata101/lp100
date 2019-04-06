package main

import (
    "./cabocha"
    "fmt"
)

func main() {
    sentenses := cabocha.LoadAsChunk("../neko.txt.cabocha")
    sentense := sentenses[7]
    for _, chunk := range sentense {
        fmt.Print(cabocha.GetChunkSurface(chunk))

        if chunk.Dst >= 0 {
            dstChunk := sentense[chunk.Dst]
            fmt.Print(" " + cabocha.GetChunkSurface(dstChunk))
        }
        fmt.Println()
    }
}
