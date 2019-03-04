package main

import (
    "os"
    "bufio"
)

func main() {
    f1, _ := os.Open("col1.txt")
    s1 := bufio.NewScanner(f1)
    col1 := []string{}
    for s1.Scan() {
        col1 = append(col1, s1.Text())
    }
    defer f1.Close()

    f2, _ := os.Open("col2.txt")
    s2 := bufio.NewScanner(f2)
    col2 := []string{}
    for s2.Scan() {
        col2 = append(col2, s2.Text())
    }
    defer f2.Close()

    f3, _ := os.OpenFile("col3.txt", os.O_WRONLY | os.O_CREATE, 0644)
    w3 := bufio.NewWriter(f3)
    for i := range col1 {
        w3.WriteString(col1[i] + "\t" + col2[i] + "\n")
    }
    w3.Flush()
    defer f3.Close()
}
