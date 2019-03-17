package morpheme

import (
    "os"
    "bufio"
    "strings"
)

type Morpheme struct {
    Surface string
    Base string
    Pos string
    Pos1 string
}

type Statement []Morpheme

func Load(filepath string) []Statement {
    f, _ := os.Open(filepath)
    defer f.Close()

    states := []Statement{}
    state := []Morpheme{}
    
    s := bufio.NewScanner(f)
    for s.Scan() {
        word := s.Text()
        if word == "EOS" {
            states = append(states, state)
            state = []Morpheme{}
            continue
        }

        results := strings.Split(word, "\t")
        list := strings.Split(results[1], ",")
        state = append(state, Morpheme{results[0], list[6], list[0], list[1]})
    }
    return states
}
