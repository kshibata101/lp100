package main

import (
	"./cabocha"
	"fmt"
)

func main() {
	sentences := cabocha.LoadAsChunk("../neko.txt.cabocha")
	for _, sentence := range sentences {
		for _, chunk := range sentence {
			for _, morph := range chunk.Morphs {
				if morph.Pos == "名詞" {
					showTree(chunk, sentence)
					break
				}
			}
		}
	}
}

func showTree(chunk cabocha.Chunk, chunks []cabocha.Chunk) {
	fmt.Print(cabocha.GetChunkSurface(chunk))
	if chunk.Dst >= 0 {
		fmt.Print(" -> ")
		showTree(chunks[chunk.Dst], chunks)
	} else {
		fmt.Println()
	}
}