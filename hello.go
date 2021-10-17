package main

import "fmt"

func main() {
	var mySlice = make([]int, 0)
	for i := 0; i < 30; i++ {
		mySlice = append(mySlice, i)
		fmt.Println("len=", len(mySlice), " cap=", cap(mySlice))
	}

	fmt.Println(mySlice[len(mySlice)-2:])
}
