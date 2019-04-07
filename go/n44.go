package main

import (
    "./cabocha"
)

func main() {
    sentences := cabocha.LoadAsChunk("../neko.txt.cabocha")
    cabocha.OutputDotFile(sentences[10], "n44.dot") // sentence is random
}
