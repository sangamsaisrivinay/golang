package main

import ("fmt"
);


func main(){
	intArr := []int {0,1,2};
	intArr1 := []int {3,3,3, 4 ,5};
	fmt.Println(intArr)
	fmt.Println(len(intArr))
	fmt.Println(cap(intArr))
	intArr = append(intArr, intArr1...)
	fmt.Println(intArr)
	fmt.Println(len(intArr))
	fmt.Println(cap(intArr))
	intArr= append(intArr, 10)
	fmt.Println(intArr)
	fmt.Println(len(intArr))
	fmt.Println(cap(intArr))
}