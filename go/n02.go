package main

import "fmt"

func concat(s1 string, s2 string) string {
    runes1 := []rune(s1)
    runes2 := []rune(s2)
    ans := ""
    for i := 0; i < len(runes1); i++ {
        ans += string(runes1[i]) + string(runes2[i])
    }
    return ans
}

func main() {
    s1 := "パトカー"
    s2 := "タクシー"
    fmt.Println(concat(s1, s2))
}
