package main 

import (
	"fmt"
	"math"
)

func main() {
	var x, y  = 3, 4.2
	var zf float64 = math.Sqrt(float64(float64(x * x) + y * y))
	var zi uint = uint(zf)
	fmt.Println(zf, zi)
}