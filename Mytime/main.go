package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("welcome to time study ogo flang")

	presentTime := time.Now()
	fmt.Println(presentTime)


	   fmt.Println(presentTime.Format("01-02-2006"))
}