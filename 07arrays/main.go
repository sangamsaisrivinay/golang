package main

import "fmt"

func main() {

	var arr [4]int
	arr[0] = 1
	arr[1] = 1
	arr[2] = 1

	fmt.Println("len: ", len(arr))
	fmt.Println(arr)

	arr1 := [5]int{1,2,3,4}

	fmt.Println("len: ", len(arr1))
	fmt.Println(arr1)

}


// len:  4
// [1 1 1 0]

// len:  5
// [1 2 3 4 0]