package main

import (
	"fmt"
	"sort"
)

func main() {

	arr1 := []int{0, 1, 2, 4,5,6,7,7}
	fmt.Println(arr1)

	fmt.Println("len: ", len(arr1))

	sort.Ints(arr1)

fmt.Println(arr1)

fmt.Println(arr1[1:3])

	fmt.Println(sort.IntsAreSorted(arr1))


	//remove element

	var index int = 2

	arr1= append(arr1[:index], arr1[index+1:]...)

	fmt.Println(arr1)

}
