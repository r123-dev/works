package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("Welcome to array in golagns")

	var fruitList =[]string{"apple"}

	fruitList=append(fruitList, "peach","guava")
	fmt.Println(fruitList)

	fruitList=append(fruitList[1:2])
	fmt.Println(fruitList)


	highscores :=make([]int,4)
	highscores[0]=1;
	highscores=append(highscores, 3,5,6,7,4,3)
	sort.Ints(highscores)
	fmt.Println(highscores)


	

}