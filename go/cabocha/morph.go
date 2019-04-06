package cabocha

import (
    "os"
    "bufio"
    "strings"
)

type Morph struct {
    Surface string
    Base string
    Pos string
    Pos1 string
}

type Sentense []Morph

func Load(filepath string) []Sentense {
    f, _ := os.Open(filepath)
    defer f.Close()

    scanner := bufio.NewScanner(f)
    sentenses := []Sentense{}
    sentense := Sentense{}
    for scanner.Scan() {
        row := scanner.Text()
        if row == "EOS" {
            if len(sentense) > 0 {
                sentenses = append(sentenses, sentense)
            }
            sentense = Sentense{}
            continue
        }

        if strings.Index(row, "*") == 0 {
            continue
        }

        words := strings.Split(row, "\t")
        analysis := strings.Split(words[1], ",")
        morph := Morph{words[0], analysis[6], analysis[0], analysis[1]}
        sentense = append(sentense, morph)
    }
    if len(sentense) > 0 {
        sentenses = append(sentenses, sentense)
    }
    return sentenses
}
