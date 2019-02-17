package main

import (
    "fmt"
    "strings"
)

func nukidashi(s string) map[string]int {
    first_indice := []int{1, 5, 6, 7, 8, 9, 15, 16, 19}
    list := strings.Split(s, " ")
    mp := map[string]int{}
    for i, word := range list {
        var is_first bool
        for _, index := range first_indice {
            if index == i+1 {
                is_first = true
            }
        }
        if is_first {
            mp[word[0:1]] = i+1
        } else {
            mp[word[0:2]] = i+1
        }
    }
    return mp
}

func main() {
    s := "Hi He Lied Because Boron Could Not Oxidize Fluorine. New Nations Might Also Sign Peace Security Clause. Arthur King Can."
    fmt.Println(nukidashi(s))
}
