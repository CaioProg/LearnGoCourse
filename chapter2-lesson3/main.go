package main

import "fmt"

func main() {
	var x = 10
	y := "Hello!!!"
	z := 10.0

	fmt.Printf("x: %v, type of %T\n", x, x)
	fmt.Printf("y: %v, type of %T\n", y, y)
	fmt.Printf("z: %v, type of %T\n", z, z)

	x = 20
	fmt.Printf("x: %v, type of %T\n", x, x)

	b := 10 < 20
	fmt.Printf("b: %v, type of %T\n", b, b)
	b = false
	fmt.Printf("b: %v, type of %T\n", b, b)
}
