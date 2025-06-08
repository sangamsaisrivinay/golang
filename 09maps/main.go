package main

import "fmt"

func main() {

	map1 := make(map[int]string)

	map1[1] = "one"
	map1[2] = "two"
	map1[3] = "three"

	fmt.Println(map1)
	fmt.Printf("map1 - %T", map1)

	for k, v := range map1 {
		fmt.Println("key - %v, value - %v", k, v)
	}

	delete(map1, 2)
	fmt.Println(map1)

}
