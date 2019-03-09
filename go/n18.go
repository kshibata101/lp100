package main

import (
    "os"
    "bufio"
    "strings"
    "strconv"
    "sort"
    "fmt"
)

type Temp struct {
    Prefecture string
    City string
    Temperature float64
    date string
}

type TempList []Temp

func (t TempList) Len() int {
    return len(t)
}

func (t TempList) Swap(i int, j int) {
    t[i], t[j] = t[j], t[i]
}

func (t TempList) Less(i, j int) bool {
    return t[i].Temperature > t[j].Temperature
}

func main() {
    filepath := "../hightemp.txt"
    f, _ := os.Open(filepath)
    s := bufio.NewScanner(f)
    var templist TempList = []Temp{}
    for s.Scan() {
        data := strings.Split(s.Text(), "\t")
        temp, _ := strconv.ParseFloat(data[2], 32)
        templist = append(templist, Temp{Prefecture: data[0], City: data[1], Temperature: temp, date: data[3]})
    }
    sort.Sort(templist)
    for _, temp := range templist {
        fmt.Println(temp.Prefecture + "\t" + temp.City + "\t" + strconv.FormatFloat(temp.Temperature, 'f', 1, 32) + "\t" + temp.date)
    }
}
