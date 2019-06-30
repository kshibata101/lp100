package main

import (
	"./cabocha"
	"fmt"
	"strings"
	"sort"
)

func main() {
	sentences := cabocha.LoadAsChunk("../neko.txt.cabocha")
	for _, sentence := range sentences {
		for _, chunk := range sentence {
			for i, morph := range chunk.Morphs {
				if !(morph.Pos == "名詞" && morph.Pos1 == "サ変接続") {
					continue
				}
				if i+1 >= len(chunk.Morphs) {
					continue
				}

				nextMorph := chunk.Morphs[i+1]
				if !(nextMorph.Pos == "助詞" && nextMorph.Base == "を") {
					continue
				}

				dstChunk := sentence[chunk.Dst]
				jutsugoVerb := ""
				for _, kMorph := range dstChunk.Morphs {
					if kMorph.Pos == "動詞" {
						jutsugoVerb = kMorph.Base
						break
					}
				}
				if jutsugoVerb == "" {
					continue
				}
				jutsugo := morph.Base + nextMorph.Base + jutsugoVerb

				// 述語(chunk,dstChunk)にかかっている助詞（とその文節）のチェック
				var srcChunks JoshiChunkList = []cabocha.Chunk{}
				for _, j := range chunk.Srcs {
					srcChunk := sentence[j]
					for _, srcMorph := range srcChunk.Morphs {
						if srcMorph.Pos == "助詞" {
							srcChunks = append(srcChunks, srcChunk)
							break
						}
					}
				}
				for _, j := range dstChunk.Srcs {
					srcChunk := sentence[j]
					for _, srcMorph := range srcChunk.Morphs {
						if srcMorph.Pos == "助詞" {
							srcChunks = append(srcChunks, srcChunk)
							break
						}
					}
				}

				// output
				fmt.Print(jutsugo)
				if len(srcChunks) > 0 {
					sort.Sort(srcChunks) // 辞書順に並べて表示するためソート
					joshi := []string{}
					bunsetsu := []string{}
					for _, srcChunk := range srcChunks {
						for _, srcMorph := range srcChunk.Morphs {
							if srcMorph.Pos == "助詞" {
								joshi = append(joshi, srcMorph.Base)
							}
						}
						bunsetsu = append(bunsetsu, cabocha.GetChunkSurface(srcChunk))
					}
					fmt.Printf("\t%s\t%s", strings.Join(joshi, " "), strings.Join(bunsetsu, " "))
				}
				fmt.Println()
			}
		}
	}
}

type JoshiChunkList []cabocha.Chunk

func (jcl JoshiChunkList) Len() int { return len(jcl) }
func (jcl JoshiChunkList) Swap(i,j int) { jcl[i], jcl[j] = jcl[j], jcl[i] }
func (jcl JoshiChunkList) Less(i,j int) bool {
	is, js := "", ""
	for _, morph := range jcl[i].Morphs {
		if morph.Pos == "助詞" {
			is = morph.Base
			break
		}
	}
	for _, morph := range jcl[j].Morphs {
		if morph.Pos == "助詞" {
			js = morph.Base
			break
		}
	}
	return is < js
}
