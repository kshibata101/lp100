package main

import (
    "os"
    "bufio"
    "regexp"
    "fmt"
)

func main() {
    r := regexp.MustCompile(`^\[\[Category:.*\]\]$`)
    
    filepath := "jawiki-uk.txt"
    f, _ := os.Open(filepath)
    s := bufio.NewScanner(f)
    for s.Scan() {
        line := s.Text()
        if r.MatchString(line) {
            fmt.Println(line)
        }
    }
}
