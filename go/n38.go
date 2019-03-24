package main

import (
    "./morpheme"
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
    values := make(plotter.Values, len(counts))
    for _, v := range counts {
        values = append(values, float64(v))
    }

    p, _ := plot.New()
    p.Title.Text = "Histogram"
    p.X.Label.Text = "Appear num"
    p.Y.Label.Text = "Word num"

    h, _ := plotter.NewHist(values, 100)
    p.Add(h)

    if err := p.Save(4 * vg.Inch, 4 * vg.Inch, "n38.png"); err != nil {
        panic(err)
    }
}
