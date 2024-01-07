package main


import "fmt"

func main() {
 var s []int
 appendSlice(&s)
 s = append(s[:10], s[len(s)-10:]...)
 fmt.Println(s)
 reverseSlice(s)
 fmt.Println(s)
}

func appendSlice(s *[]int) {
 for i := 0; i < 100; i++ {
  *s = append(*s, (i + 1))
 }
}

func reverseSlice(s []int) {
 for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
  s[i], s[j] = s[j], s[i]
 }
}