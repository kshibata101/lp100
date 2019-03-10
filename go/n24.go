package main

import (
    "regexp"
    "os"
    "bufio"
    "fmt"
)

func main() {
    r := regexp.MustCompile(`\[\[(File|ファイル):(.*?)\|`)

    filepath := "jawiki-uk.txt"
    f, _ := os.Open(filepath)
    defer f.Close()

    s := bufio.NewScanner(f)
    for s.Scan() {
        res := r.FindStringSubmatch(s.Text())
        if res != nil {
            fmt.Println(res[2])
        }
    }
}
