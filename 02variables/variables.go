package main

import "fmt"

//global decalration
const pi = 3
const login = "yes"
var wish = "hi"

//invalid use of walrus outside function
// s6:="hi"

func main() {

	var num int = 5
	var str string = "hi"
	var bl bool = true

	fmt.Println("vale of num ", num)
	fmt.Printf("type of num %T \n", num)

	fmt.Println("vale of str ", str)
	fmt.Printf("type of str %T \n", str)

	fmt.Println("vale of bl ", bl)
	fmt.Printf("type of bl %T \n", bl)

	//implicit conversion
	var n1 = "hi"
	fmt.Println("vale of n1 ", n1)
	fmt.Printf("type of n1 %T \n", n1)

	//walrus opeerator
	rand:=10
	fmt.Println("vale of rand ", rand)
	fmt.Printf("type of rand %T \n", rand)

	fmt.Println("vale of pi ", pi)
	fmt.Printf("type of pi %T \n", pi)

	fmt.Println("vale of login ", login)
	fmt.Printf("type of login %T \n", login)

	fmt.Println("vale of wish ", wish)
	fmt.Printf("type of wish %T \n", wish)

	//default values
	var s1 int
	var s2 string
	var s3 bool
	var s4 float32
	var s5 float64

	fmt.Printf("s1: type = %T, value = %v\n", s1, s1)
	fmt.Printf("s2: type = %T, value = %v\n", s2, s2)
	fmt.Printf("s3: type = %T, value = %v\n", s3, s3)
	fmt.Printf("s4: type = %T, value = %v\n", s4, s4)
	fmt.Printf("s5: type = %T, value = %v\n", s5, s5)


}
