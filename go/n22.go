package main

import (
    "regexp"
    "os"
    "bufio"
    "fmt"
)

func main() {
    r := regexp.MustCompile(`^\[\[Category:(.*)\]\]$`)
    
    filepath := "jawiki-uk.txt"
    f, _ := os.Open(filepath)
    defer f.Close()

    s := bufio.NewScanner(f)
    for s.Scan() {
        l := s.Text()
        res := r.FindStringSubmatch(l)
        if res != nil {
            fmt.Println(res[1])
        }
    }
}
