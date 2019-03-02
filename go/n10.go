package main

import (
    "fmt"
    "os"
    "bufio"
)

func line_count(filepath string) int {
    f, err := os.Open(filepath)
    if err != nil {
        return 0
    }

    s := bufio.NewScanner(f)
    linenum := 0
    for s.Scan() {
        linenum += 1    
    }
    return linenum
}

func main() {
    n := line_count("../hightemp.txt")
    fmt.Println(fmt.Sprint(n))
}
