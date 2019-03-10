package main

import (
    "os"
    "bufio"
    "regexp"
    "fmt"
    "strings"
)

func main() {
    texts := []string{}

    r := regexp.MustCompile(`^\|(.*)`)
    rs := regexp.MustCompile(`^\{\{基礎情報`)
    re := regexp.MustCompile(`^\}\}`)
    
    filepath := "jawiki-uk.txt"
    f, _ := os.Open(filepath)
    defer f.Close()

    basicInfo := false
    s := bufio.NewScanner(f)
    for s.Scan() {
        l := s.Text()
        
        if re.Match([]byte(l)) {
            basicInfo = false
        }
        if basicInfo {
            res := r.FindStringSubmatch(l)
            if res == nil {
                texts[len(texts)-1] += l
            } else {
                texts = append(texts, res[1])
            }
        }
        if rs.Match([]byte(l)) {
            basicInfo = true
        }
    }

    rst := regexp.MustCompile(`'{2,}`)
    m := make(map[string]string)
    for _, text := range texts {
        res := strings.Split(text, " = ")
        if (len(res) == 2) {
            m[res[0]] = rst.ReplaceAllString(res[1], "")
        }
    }

    fmt.Println(m)
}
