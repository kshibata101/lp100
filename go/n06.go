package main

import (
    "fmt"
)

func bigram(s string) []string {
    chars := []rune(s)
    list := []string{}
    for i := 0; i < len(chars) - 1; i++ {
        list = append(list, string(chars[i]) + string(chars[i+1]))
    }
    return list
}

func printSet(title string, set map[string]struct{}) {
    fmt.Println(title)
    for key, _ := range set {
        fmt.Print(key)
        fmt.Print(" ")
    }
    fmt.Println()
}

func main() {
    s1 := "paraparaparadise"
    s2 := "paragraph"

    x_bigram := bigram(s1)
    y_bigram := bigram(s2)

    x := make(map[string]struct{})
    y := make(map[string]struct{})

    for _, val := range x_bigram {
        x[val] = struct{}{}
    }
    for _, val := range y_bigram {
        y[val] = struct{}{}
    }

    // 集合
    printSet("X", x)
    printSet("Y", y)

    x_plus_y := make(map[string]struct{}) // 和集合
    x_by_y := make(map[string]struct{}) // 積集合
    x_minus_y := make(map[string]struct{}) // 差集合

    for key, _ := range x {
        x_plus_y[key] = struct{}{}
        if _, exist := y[key]; exist {
            x_by_y[key] = struct{}{}
        } else {
            x_minus_y[key] = struct{}{}
        }
    }
    for key, _ := range y {
        x_plus_y[key] = struct{}{}
        if _, exist := x[key]; !exist {
            x_minus_y[key] = struct{}{}
        }
    }
    printSet("X+Y", x_plus_y)
    printSet("X*Y", x_by_y)
    printSet("X-Y", x_minus_y)

    if _, exist := x["se"]; exist {
        fmt.Println("X includes 'se'")
    } else {
        fmt.Println("X does not include 'se'")
    }
    if _, exist := y["se"]; exist {
        fmt.Println("Y includes 'se'")
    } else {
        fmt.Println("Y does not include 'se'")
    }
}
