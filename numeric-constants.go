package main 

const (
	Big = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int {
	return x * 10 + 1
}

func main() {
 println(needInt(Small))
}