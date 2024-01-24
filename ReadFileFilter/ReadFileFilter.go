package main

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
    "time"
)

func main() {
    timeCall_1 := metricTimeCall(PrintAllFilesWithFilter)
    timeCall_2 := metricTimeCall(PrintAllFilesWithFilterClosure)
    timeCall_1("C:\\Users\\Nupgod\\Desktop\\Yandex_GO", ".go")
    fmt.Println("Closer Func")
    timeCall_2("C:\\Users\\Nupgod\\Desktop\\Yandex_GO", ".go")
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
// Поиск файла по фильтру через замыкание
func PrintAllFilesWithFilterClosure(path string, filter string) {
    // создаём переменную, содержащую функцию обхода
    // мы создаём её заранее, а не через оператор :=, чтобы замыкание могло сослаться на него
    var walk func(string)
    walk = func(path string) {
        // получаем список всех элементов в папке (и файлов, и директорий)
        files, err := os.ReadDir(path)
        if err != nil {
            fmt.Println("unable to get list of files", err)
            return
        }
        //  проходим по списку
        for _, f := range files {
            // получаем имя элемента
            // filepath.Join — функция, которая собирает путь к элементу с разделителями
            filename := filepath.Join(path, f.Name())
            // печатаем имя элемента, если путь к нему содержит filter, который получим из внешнего контекста
            if strings.Contains(filename, filter) {
                fmt.Println(filename)
            }
            // если элемент — директория, то вызываем для него рекурсивно ту же функцию
            if f.IsDir() {
                walk(filename)
            }
        }
    }
    // теперь вызовем функцию walk
    walk(path)
}

func metricTimeCall(f func(string, string)) func(string, string) {
    return func(s1 string, s2 string) {
        start := time.Now() // получаем текущее время
        f(s1, s2)
        fmt.Println("Execution time: ", time.Now().Sub(start)) // получаем интервал времени как разницу между двумя временными метками
    }
}
