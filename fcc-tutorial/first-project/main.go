package main

import ("fmt"
"unicode/utf8"
);


func main(){
	fmt.Println("c"+"d")
	fmt.Println(utf8.RuneCountInString("α"))

}