package main

import (
    "./morpheme"
    "sort"
    "math"
    "gonum.org/v1/plot"
    "gonum.org/v1/plot/plotter"
    "gonum.org/v1/plot/vg"
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

    p, err := plot.New()
    if err != nil {
        panic(err)
    }

    p.Title.Text = "Log-Log Graph"
    p.X.Label.Text = "Log(Rank)"
    p.Y.Label.Text = "Log(Frequency)"
    p.Add(plotter.NewGrid())

    xys := make(plotter.XYs, len(list))
    for i := range xys {
        xys[i].X = math.Log(float64(i+1))
        xys[i].Y = math.Log(float64(list[i].Value))
    }

    s, err := plotter.NewScatter(xys)
    if err != nil {
        panic(err)
    }
    
    p.Add(s)
    p.Legend.Add("scatter", s)

    if err := p.Save(4*vg.Inch, 4*vg.Inch, "n39.png"); err != nil {
        panic(err)
    }
}
