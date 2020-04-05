package main

import "fmt"

func main() {
	sum := math.Sum([]int{5, 8, -3})

	if sum != 10 {
		msg := fmt.Sprintf("Expected 10, got %d", sum)
		panic(msg)
	}
	fmt.Println("Passed!")
}
