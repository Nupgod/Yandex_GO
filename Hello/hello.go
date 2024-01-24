package main

import (
    "fmt"
    "time"
)

func main() {
    date := time.Now()

    switch y:= date.Year(); {
    case y > 1946 && y < 1965:
        fmt.Println("Hello, Boomer!")
    case y >= 1965 && y < 1981:
        fmt.Println("Hello, X!")
    case y >= 1981 && y < 1996:
        fmt.Println("Hello, Millenial!")
    case y >= 1997 && y < 2012:
        fmt.Println("Hello, Zoomer!")
    case y >= 2013:
        fmt.Println("Hello, Alpha!")
    default:
        fmt.Println("Hello!")
    }
}