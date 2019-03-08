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

    s := bufio.NewScanner(f)
    list := []string{}
    for s.Scan() {
        list = append(list, s.Text())
    }
    for i := 0; i < n; i++ {
        fmt.Println(list[len(list) - n + i])
    }
}
