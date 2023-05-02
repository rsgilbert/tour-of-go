package main 

import "fmt"

func add(x int, y float64) float64 {
	return float64(x) + y;
}

func addThree(a, b, c int) int {
	return a + b + c;
}

func main() {
	fmt.Println(add(100, 200.3))
	fmt.Println(addThree(10, 20, 8))
}