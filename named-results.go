package main 

import "fmt"

func split(sum int) (x, y int) {
	x = (sum / 2) - 5;
	y = sum - x;
	return;
}

func main() {
	fmt.Println(split(51))
}