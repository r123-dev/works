package main

import "fmt"

func main() {

	fmt.Println("Structs in golang")

	hitesh := User{"Hitesh" , "hitesh@go.dev",true,16}
   fmt.Println(hitesh)
}


type User struct {

	Name string
	Email string
	Status bool
	Age int

}
