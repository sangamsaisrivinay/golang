package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {

	reader:= bufio.NewReader(os.Stdin)

	fmt.Println("please enter a number:")
	input, err := reader.ReadString('\n')

	if(nil!=err){
		fmt.Println("error: ", err)
	}
	fmt.Print(input)

	fmt.Printf("type of input: %T", input)

}