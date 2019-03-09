package main

import (
    "os"
    "bufio"
    "encoding/json"
    "strings"
    "fmt"
)

type Wiki struct {
    Text string
    Title string
}

func main() {
    filepath := "../jawiki-country.json"
    f, _ := os.Open(filepath)
    defer f.Close()

    var wikis []Wiki
    s := bufio.NewScanner(f)
    for s.Scan() {
        var wiki Wiki
        in := s.Bytes()
        json.Unmarshal(in, &wiki);

        if strings.Contains(wiki.Text, "イギリス") {
            fmt.Println(string(in))
            wikis = append(wikis, wiki)
        }
    }

}
