package main

import (
    "os"
    "bufio"
    "strings"
    "fmt"
)

func main() {
    filepath := "../hightemp.txt"
    f, err := os.Open(filepath)
    if err != nil {
        fmt.Println(err)
    }
    
    s := bufio.NewScanner(f)
    set := make(map[string]struct{})
    for s.Scan() {
        text := strings.Split(s.Text(), "\t")[0]
        set[text] = struct{}{}
    }
    for k, _ := range set {
        fmt.Println(k)
    }
}
