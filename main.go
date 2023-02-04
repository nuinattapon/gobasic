package main

import (
	"fmt"
)

var x int64 = 1_234

func main() {
	/*
	 * main function
	 */

	fmt.Println("Hello, world!")
	s := "Golang"
	i := 1_234
	f := 1_234.56
	fmt.Printf("%T %#v\n", s, s)
	fmt.Printf("%T %#v\n", i, i)
	fmt.Printf("%T %#v\n", f, f)

	fmt.Printf("%T %#v\n", x, x)

}
