package main

import (
	"euler/natural"
	"fmt"
)

func main() {
	sum := 0
	for i := range natural.Fibo() {
		if i >= 4000000 {
			break
		}
		if i%2 == 0 {
			sum += i
		}
	}
	fmt.Println(sum)
}
