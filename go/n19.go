package main

import (
    "os"
    "bufio"
    "strings"
    "fmt"
    "sort"
)

type Pair struct {
    Key string
    Value int
}

type PairList []Pair

func (pl PairList) Len() int {
    return len(pl)
}

func (pl PairList) Swap(i, j int) {
    pl[i], pl[j] = pl[j], pl[i]
}

func (pl PairList) Less(i, j int) bool {
    return pl[i].Value > pl[j].Value
}

func main() {
    hash := make(map[string]int)
    
    filepath := "../hightemp.txt"
    f, _ := os.Open(filepath)
    s := bufio.NewScanner(f)
    for s.Scan() {
        pref := strings.Split(s.Text(), "\t")[0]
        hash[pref] += 1
    }

    var pairlist PairList = []Pair{}
    for k, v := range hash {
        pairlist = append(pairlist, Pair{Key: k, Value: v})
    }
    sort.Sort(pairlist)
    for _, pair := range pairlist {
        fmt.Println(pair.Key)
    }
}
