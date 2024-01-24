package main

import "fmt"

func main() {
	for i:= 1; i <= 15; i++ {
	if i % 3 == 0 {
		if i % 5 == 0 {
			fmt.Println("FizzBuzz")
			continue;
		}
		fmt.Println("Fizz")
	} else if i % 5 == 0 {
		fmt.Println("Buzz")
	} else {
		fmt.Printf("%d\n", i)
	}
  }
}