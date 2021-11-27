package main

import "fmt"

type Point struct {
	Lat    float64
	Long   float64
	Heigth float64
}

func main() {
	x := []int{1, 2, 3, 4, 5}

	fmt.Println(x)

	func(mySlice []int) {
		mySlice[0] = 10
	}(x)

	fmt.Println(x)

}
