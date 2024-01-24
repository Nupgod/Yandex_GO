package main
// Пример взят из стандартной библиотеки
    import (
		"bytes"
		"fmt"
	)
	
	func main() {
		bSlice := []byte(" \t\n a lone gopher \n\t\r\n")
		fmt.Printf("%s", bytes.TrimSpace(bSlice)) // a lone gopher
		fmt.Printf("%s", bSlice)  // \t\n a lone gopher \n\t\r\n 
		
	}