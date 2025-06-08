package main

import "fmt"

func main()  {
	var ptr *int;

	x := 10

	xadd := &x

	fmt.Println("ptr: ", ptr)
	fmt.Println("x-add: ", xadd)
	fmt.Println("*xadd: ", *xadd)

	*xadd= *xadd + 3

	fmt.Println("x: ", x)
	fmt.Println("*xadd: ", *xadd)

}