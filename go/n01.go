package main

import "fmt"

func toridashi(s string) string {
    runes := []rune(s)
    arr := [4]int{0, 2, 4, 6}
    ans := ""
    for i := 0; i < len(arr); i++ {
        ans += string(runes[arr[i]])
    }
    return ans
}

func main() {
    s := "パタトクカシーー"
    fmt.Println(toridashi(s))
}
