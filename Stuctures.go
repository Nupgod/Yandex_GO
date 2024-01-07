package main

import (
	"fmt"
	"time"
	"encoding/json"
)

type Person struct {
    Name        string `json:"name"`
    Email       string `json:"email"`
    DateOfBirth time.Time `json:"date"`
}

func main() {
	personInfo := Person {Name: "Aлекс", Email: "alex@yandex.ru"} 
}