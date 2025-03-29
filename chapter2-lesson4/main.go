package main

var y = 7

func main() {
	y := 10

	testfunc(y)
}

func testfunc(x int) {
	println(x)
	println(y)
}
