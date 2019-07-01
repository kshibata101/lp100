package cabocha

import (
    "os"
    "bufio"
    "strings"
    "strconv"
    "regexp"
)

type Morph struct {
    Surface string
    Base string
    Pos string
    Pos1 string
}

type Chunk struct {
    Morphs []Morph
    Dst int
    Srcs []int
}

type Sentense []Chunk

func LoadAsMorph(filepath string) [][]Morph {
    sentenses := [][]Morph{}
    sentense := []Morph{}

    f, _ := os.Open(filepath)
    defer f.Close()

    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        row := scanner.Text()
        if row == "EOS" {
            if len(sentense) > 0 {
                sentenses = append(sentenses, sentense)
            }
            sentense = []Morph{}
            continue
        }

        if strings.Index(row, "*") == 0 {
            continue
        }

        words := strings.Split(row, "\t")
        morph_texts := strings.Split(words[1], ",")
        morph := Morph{words[0], morph_texts[6], morph_texts[0], morph_texts[1]}
        sentense = append(sentense, morph)
    }
    return sentenses
}

func LoadAsChunk(filepath string) []Sentense {
    regex := regexp.MustCompile(`^\* ([0-9]+) (-?[0-9]+)D`)

    sentenses := []Sentense{}
    sentense := Sentense{}
    chunk := Chunk{[]Morph{}, 0, []int{}}

    f, _ := os.Open(filepath)
    defer f.Close()

    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        row := scanner.Text()
        if row == "EOS" {
            if len(chunk.Morphs) > 0 {
                sentense = append(sentense, chunk)
            }
            if len(sentense) > 0 {
                sentenses = append(sentenses, sentense)
            }
            sentense = Sentense{}
            chunk = Chunk{[]Morph{}, 0, []int{}}
            continue
        }

        if strings.Index(row, "*") == 0 {
            if len(chunk.Morphs) > 0 {
                sentense = append(sentense, chunk)
            }

            // 第一引数
            morphs := []Morph{}

            // 第二引数
            result := regex.FindStringSubmatch(row)
            dst, _ := strconv.Atoi(result[2])
            currentIndex, _ := strconv.Atoi(result[1])

            // 第三引数
            srcs := []int{}
            for i, c := range sentense {
                if c.Dst == currentIndex {
                    srcs = append(srcs, i)
                }
            }

            chunk = Chunk{morphs, dst, srcs}
            continue
        }

        words := strings.Split(row, "\t")
        morph_texts := strings.Split(words[1], ",")
        morph := Morph{words[0], morph_texts[6], morph_texts[0], morph_texts[1]}
        chunk.Morphs = append(chunk.Morphs, morph)
    }
    return sentenses
}

func GetChunkSurface(chunk Chunk) string {
    text := ""
    for _, morph := range chunk.Morphs {
        text += morph.Surface
    }
    return text
}

func (c Chunk) GetChunkSurface() string {
    text := ""
    for _, morph := range c.Morphs {
        text += morph.Surface
    }
    return text
}

func GetChunkSurfaceWithoutSymbol(chunk Chunk) string {
    text := ""
    for _, morph := range chunk.Morphs {
        if morph.Pos != "記号" {
            text += morph.Surface
        }
    }
    return text
}

func HasPos(chunk Chunk, pos string) bool {
    for _, morph := range chunk.Morphs {
        if morph.Pos == pos {
            return true
        }
    }
    return false
}

func OutputDotFile(sentence Sentense, filepath string) {
    f, err := os.Create(filepath)
    if err != nil {
        panic(err)
    }
    defer f.Close()

    w := bufio.NewWriter(f)
    w.WriteString("digraph graphname {\n")

    for _, chunk := range sentence {
        if chunk.Dst >= 0 {
            w.WriteString("    ")
            w.WriteString(GetChunkSurfaceWithoutSymbol(chunk))
            w.WriteString(" -> ")
            w.WriteString(GetChunkSurfaceWithoutSymbol(sentence[chunk.Dst]))
            w.WriteString(";\n")
        }
    }

    w.WriteString("}\n")
    w.Flush()
}
