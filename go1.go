package main

import (
	"fmt"
)

func main() {
	var str = "Grüezi!😂👉"
	for _, r := range str {
		fmt.Printf("%s", string(r))

	}

	teams := map[string][]string{
		"First":  []string{"1", "2", "3"},
		"Second": []string{"1", "2", "3"},
	}
	fmt.Println(teams)
	myMap := map[int]string{}
	myMap[1] = "Frist!"

	if _, ok := myMap[2]; !ok {
		fmt.Println("myMap[2] doesn't exists")
	}
}
