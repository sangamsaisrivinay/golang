package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()

	fmt.Println("now - ", now)
	fmt.Println("TIME: ", now.Format("05:15:56 Monday Jan 4"))

}
