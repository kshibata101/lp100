package main

import (
    "os"
    "bufio"
    "strings"
)

func save_column_1_2(filepath string) {
    f, err := os.Open(filepath)
    if err != nil {
        return
    }
    defer f.Close()

    s := bufio.NewScanner(f)
    col1 := []string{}
    col2 := []string{}
    for s.Scan() {
        words := strings.Split(s.Text(), "\t")
        col1 = append(col1, words[0])
        col2 = append(col2, words[1])
    }

    f1, _ := os.OpenFile("./col1.txt", os.O_WRONLY | os.O_CREATE, 0644)
    w1 := bufio.NewWriter(f1)
    w1.WriteString(strings.Join(col1, "\n"))
    w1.Flush()
    defer f1.Close()
    
    f2, _ := os.OpenFile("./col2.txt", os.O_WRONLY | os.O_CREATE, 0644)
    w2 := bufio.NewWriter(f2)
    w2.WriteString(strings.Join(col2, "\n"))
    w2.Flush()
    defer f2.Close()
}

func main() {
   save_column_1_2("../hightemp.txt")
}
