package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)



func main(){

	fmt.Println("Welcome to our pizza app")

	reader :=bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	  fmt.Println("thanks for rating ,", input)


	  numRating,err:=strconv.ParseFloat(strings.TrimSpace(input),64)

      if err !=nil {
		fmt.Println(err)
		
	  } else {
		fmt.Println("Added 1 to you rating : ",numRating+1)
	  }


}