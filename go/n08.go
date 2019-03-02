package main

import (
    "fmt"
)

func cipher(s string) string {
    chars := []rune(s)
    str := ""
    for i := 0; i < len(chars); i++ {
        if 'a' <= chars[i] && chars[i] <= 'z' {
            str += string(219 - chars[i])
        } else {
            str += string(chars[i])
        }
    }
    return str
}

func main() {
    s := "abcxyzABCXYZ"
    fmt.Println(cipher(s))
}
