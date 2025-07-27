package main

import(
	"fmt"
)

func main(){

	var i int32
	var p *int32 = new (int32)
	i = 10
	p = &i
	fmt.Println(i)
	fmt.Println(*p)
	fmt.Println(p)
	fmt.Println(&i)
	*p = 15
	fmt.Println(*p)
	fmt.Println(i)
	fmt.Println(p)
	fmt.Println(&i)
}