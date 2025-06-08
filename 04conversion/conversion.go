package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main()  {

	reader:= bufio.NewReader(os.Stdin)
	fmt.Println("enter a number: ")
	input, _ := reader.ReadString('\n')

	fmt.Println("inp: ", input)

	rating, err:= strconv.ParseFloat(strings.TrimSpace(input), 64)

	if(nil!=err){
		fmt.Println("err: ", err)
	}

	fmt.Println("rating", rating+1)
	newRating:=rating+1


	fmt.Println(newRating)

}