package main

import (
	"fmt"
	"runtime"
)

func _main() {
	fmt.Println(runtime.GOMAXPROCS(20))
	fmt.Println(runtime.NumCPU())
}
