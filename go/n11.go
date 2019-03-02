package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
)

func replace(filepath string) string {
    f, err := os.Open(filepath)
    if err != nil {
        return ""
    }

    s := bufio.NewScanner(f)
    lines := []string{}
    for s.Scan() {
        r := strings.Replace(s.Text(), "\t", " ", -1)
        lines = append(lines, r)
    }
    return strings.Join(lines, "\n")
}

func main() {
    fmt.Println(replace("../hightemp.txt"))
}
