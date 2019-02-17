package main

import "fmt"

func get_reverse(s string) string {
    ans := ""
    for i := 0; i < len(s); i++ {
        ans += s[len(s)-i-1:len(s)-i]
    }
    return ans
}

func main() {
    res := get_reverse("stressed")
    fmt.Println(res)
}
