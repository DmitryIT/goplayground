package main

import (
	"fmt"

	"github.com/DmitryIT/goplayground/operations"
)

func main() {
	i, j := 1, 2
	fmt.Printf("%d + %d = %d\n", i, j, operations.Add(i, j))
}
