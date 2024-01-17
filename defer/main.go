package main

import "fmt"

func main() {
	defer fmt.Println("hello")
	defer fmt.Println("world")
	fmt.Println("Hello")
}