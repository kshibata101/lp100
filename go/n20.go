package main

import (
    "os"
    "bufio"
    "encoding/json"
    "strings"
    "fmt"
    "io"
)

type Wiki struct {
    Text string `json:"text"`
    Title string `json:"title"`
}

func main() {
    filepath := "../jawiki-country.json"
    f, _ := os.Open(filepath)

    r := bufio.NewReaderSize(f, 1000000)
    for {
        line, _, err := r.ReadLine()
        if err == io.EOF {
            break
        } else if err != nil {
            continue
        }
        
        var wiki Wiki
        err = json.Unmarshal(line, &wiki);
        if err != nil {
            continue
        }

        if strings.Contains(wiki.Title, "イギリス") {
            fmt.Println(string(line))
        }
    }
}
