package main

import "fmt"

func main() {
	a := 1
	p := &a
	b := &p
	
	*p = 3
	fmt.Printf("%d\n", a)
	**b = 5
	fmt.Printf("%d\n", a)
	
	a += 4 + *p + **b
	
	fmt.Printf("%d",a)
  } 