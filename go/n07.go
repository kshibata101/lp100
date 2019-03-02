package main

import (
    "fmt"
)

func generate(x int, y string, z float64) string {
    return fmt.Sprint(x) + "時の" + y + "は" + fmt.Sprint(z)
}

func main() {
    fmt.Println(generate(12, "気温", 22.4))
}
