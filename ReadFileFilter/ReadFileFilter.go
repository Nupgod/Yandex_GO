package main

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
)

func main() {
    PrintAllFilesWithFilter(".", "go")
}

func PrintAllFilesWithFilter(path string, filter string){
    files, err := os.ReadDir(path)
    if err != nil {
        fmt.Println("unable to get list of files", err)
        return
    }
    for _, f := range files {
        filename := filepath.Join(path, f.Name())
        if strings.Contains(filename, filter) {
            fmt.Println(filename)
        }
        if f.IsDir() {
            PrintAllFilesWithFilter(filename, filter)
        }
    }
}