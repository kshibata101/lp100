package main

import (
    "os"
    "bufio"
    "regexp"
    "strings"
    "fmt"
)

func main() {
    r1 := regexp.MustCompile(`^{{基礎情報`)
    r2 := regexp.MustCompile(`^}}`)
    r3 := regexp.MustCompile(`^\|(.*)`)
    
    filepath := "jawiki-uk.txt"
    f, _ := os.Open(filepath)
    defer f.Close()

    mem := false
    texts := []string{}
    s := bufio.NewScanner(f)
    for s.Scan() {
        l := s.Text()
        if r2.Match([]byte(l)) {
            mem = false
        }
        if mem {
            res := r3.FindStringSubmatch(l)
            if res == nil {
                texts[len(texts) - 1] += l
            } else {
                texts = append(texts, res[1])
            }
        }
        if r1.Match([]byte(l)) {
            mem = true
        }
    }
    m := make(map[string]string)
    for _, text := range texts {
        res := strings.Split(text, " = ")
        if len(res) == 2 {
            m[res[0]] = res[1]
        }
    }

    fmt.Println(m)
}
