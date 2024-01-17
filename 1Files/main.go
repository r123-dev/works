package main

import "fmt"

func main() {

	var arr = []int{1, 2, 3, 1, 4, 5, 2, 5, 2, 4, 4, 5}
	mp := make(map[int]int)

	var k int = 6

	for i := 0; i < 12; i++ {

		if mp[k-arr[i]]>0 {
			fmt.Println(" values are present")
		} else{
			mp[arr[i]]++
		}
	}

}