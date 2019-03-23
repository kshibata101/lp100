package main

import (
    "./morpheme"
    "sort"
    "gonum.org/v1/plot"
    "gonum.org/v1/plot/plotter"
    "gonum.org/v1/plot/plotutil"
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

    keys := []string{}
    values := plotter.Values{}
    for _, entry := range list[0:10] {
        keys = append(keys, entry.Key)
        values = append(values, float64(entry.Value))
    }
    
    plo, _ := plot.New()
    plo.Title.Text = "Chart"
    plo.X.Label.Text = "Word"
    plo.Y.Label.Text = "Histgram"

    w := vg.Points(10)

    bars, _ := plotter.NewBarChart(values, w)

    bars.LineStyle.Width = vg.Length(0)
    bars.Color = plotutil.Color(3)
    bars.Offset = 0
    bars.Horizontal = false

    plo.Add(bars)
    plo.Legend.Add("values", bars)
    plo.Legend.Top = true
    plo.Legend.Left = false
    plo.NominalX(keys...)

    plo.Save(5 * vg.Inch, 3 * vg.Inch, "n37.png")
}
