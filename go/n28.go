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
    rmrk := regexp.MustCompile(`\[\[(([^\[]*)\|)?([^\[]*?)\]\]`)
    rref := regexp.MustCompile(`<ref(.*?)</ref>`)
    rref2 := regexp.MustCompile(`<ref(.*?)/>`)
    rhttp := regexp.MustCompile(`\[(.*?)\]`)
    rtpl := regexp.MustCompile(`\{\{([^\}]*\|)?([^\}]*?)\}\}`)
    rstar := regexp.MustCompile(`\*`)
    
    m := make(map[string]string)
    for _, text := range texts {
        res := strings.Split(text, " = ")
        if (len(res) == 2) {
            word := rst.ReplaceAllString(res[1], "")
            word = rmrk.ReplaceAllString(word, "$3")
            word = rref.ReplaceAllString(word, "")
            word = rref2.ReplaceAllString(word, "")
            word = rhttp.ReplaceAllString(word, "$1")
            word = rtpl.ReplaceAllString(word, "$2")
            word = rstar.ReplaceAllString(word, "")
            m[res[0]] = word
        }
    }

    for k, v := range m {
        fmt.Println(k + "\t" + v)
    }
}
