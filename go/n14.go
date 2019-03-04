package main

import (
    "fmt"
    "flag"
    "os"
    "bufio"
    "strconv"
)

func main() {
    flag.Parse()
    n, _ := strconv.Atoi(flag.Arg(0))

    filepath := "../hightemp.txt"
    f, _ := os.Open(filepath)

    s := bufio.NewScanner(f)
    i := 0
    for s.Scan() {
        i++
        if i > n {
            break
        }
        fmt.Println(s.Text())
    }
    defer f.Close()
}
