package main

import (
    "fmt"
    "flag"
    "strconv"
    "os"
    "bufio"
)

func main() {
    flag.Parse()
    n, _ := strconv.Atoi(flag.Arg(0))

    filepath := "../hightemp.txt"
    f, _ := os.Open(filepath)

    data := []string{}
    s := bufio.NewScanner(f)
    for s.Scan() {
        data = append(data, s.Text())
    }

    l := len(data) / n
    for i := 0; i < len(data); i++ {
        if i > 0 && i % l == 0 {
            fmt.Println()
        }
        fmt.Println(data[i])
    }
}
