package main

import (
    "./morpheme"
    "sort"
    "fmt"
)

type Entry struct {
    Key string
    Value int
}

type EntryList []Entry

func (list EntryList) Len() int {
    return len(list)
}

func (list EntryList) Swap(i, j int) {
    list[i], list[j] = list[j], list[i]
}

func (list EntryList) Less(i, j int) bool {
    return list[i].Value > list[j].Value
}

func main() {
    filepath := "../neko.txt.mecab"
    states := morpheme.Load(filepath)

    counts := make(map[string]int)
    for _, state := range states {
        for _, morpheme := range state {
            if _, ok := counts[morpheme.Base]; !ok {
                counts[morpheme.Base] = 0
            }
            counts[morpheme.Base] += 1
        }
    }
    list := EntryList{}
    for k, v := range counts {
        list = append(list, Entry{k, v})
    }
    sort.Sort(list)

    fmt.Println(list)
}
