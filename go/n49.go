package main

import (
	"./cabocha"
	"fmt"
)

func main() {
	sentences := cabocha.LoadAsChunk("../neko.txt.cabocha")
	for _, sentence := range sentences {
		var nounChunkIndexes []int
		for i, chunk := range sentence {
			for _, morph := range chunk.Morphs {
				if morph.Pos == "名詞" {
					nounChunkIndexes = append(nounChunkIndexes, i)
					break
				}
			}
		}
		n := len(nounChunkIndexes)
		for i := 0; i < n; i++ {
			for j := i+1; j < n; j++ {
				si, sj := nounChunkIndexes[i], nounChunkIndexes[j]
				ki, kj := si, sj
				for ki != kj {
					if ki < kj {
						ki = sentence[ki].Dst
					} else {
						kj = sentence[kj].Dst
					}
				}

				// si
				for _, morph := range sentence[si].Morphs {
					if morph.Pos == "名詞" {
						fmt.Print("X")
					} else {
						fmt.Print(morph.Surface)
					}
				}
				for l := sentence[si].Dst; l < ki; l = sentence[l].Dst {
					fmt.Printf(" -> %s", sentence[l].GetChunkSurface())
				}

				if sj == kj {
					fmt.Print(" -> ")
					for _, morph := range sentence[ki].Morphs {
						if morph.Pos == "名詞" {
							fmt.Print("Y")
						} else {
							fmt.Print(morph.Surface)
						}
					}
				} else {
					fmt.Print(" | ")
					for _, morph := range sentence[sj].Morphs {
						if morph.Pos == "名詞" {
							fmt.Print("Y")
						} else {
							fmt.Print(morph.Surface)
						}
					}
					for l := sentence[sj].Dst; l < kj; l = sentence[l].Dst {
						fmt.Printf(" -> %s", sentence[l].GetChunkSurface())
					}
					fmt.Printf(" | %s", sentence[kj].GetChunkSurface())
				}
				fmt.Println()
			}
		}
	}
}
