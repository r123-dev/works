package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){


	welcome :="Welcome to user input"
	fmt.Println(welcome);

	reader :=bufio.NewReader(os.Stdin)
	fmt.Println("enter the ratinf for your pizza:")

	// comma ok// err ok
	input, _ := reader.ReadString('\n')
	fmt.Println("thanks for rating, ",input)

}