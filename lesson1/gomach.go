package main

import "fmt"
import "math"

func main() {
	const (
		a = 3
		b = 4
	)
	c := math.Sqrt(a*a + b*b)
	fmt.Printf("%.2f", c)

}
