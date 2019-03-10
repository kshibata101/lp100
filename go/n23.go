package main

import (
    "os"
    "bufio"
    "fmt"
    "regexp"
)

func main() {
    r := regexp.MustCompile(`^=(=+)([^=]+)=+$`)

    filepath := "jawiki-uk.txt"
    f, _ := os.Open(filepath)
    defer f.Close()

    s := bufio.NewScanner(f)
    for s.Scan() {
        res := r.FindStringSubmatch(s.Text())
        if res != nil {
            fmt.Println(res[2] + "\t" + fmt.Sprint(len([]rune(res[1]))))
        }
    }
}
