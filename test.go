package main

import "fmt"

type student struct {
	id   int
	name string
}

func (t student) Doing() {
	fmt.Printf("abc")
}

func main() {

	var y int = 1
	var x int64
	x = int64(y)
	fmt.Print(x)
}
