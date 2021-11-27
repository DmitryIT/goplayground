package main

import "fmt"

type foobar interface {
	foo()
	bar()
}

type itemA struct{}

func (a *itemA) foo() {
	fmt.Println("foo() on A")
}

func (a *itemA) bar() {
	fmt.Println("bar() on A")
}

type itemB struct{}

func (b *itemB) foo() {
	fmt.Println("foo() on B")
}

func (b *itemB) bar() {
	fmt.Println("bar() on B")
}

func doFoo(item foobar) {
	item.foo()
}

type itemC struct{}
type empty interface{}

func main() {
	var _ foobar = &itemA{}
	doFoo(&itemA{})
	doFoo(&itemB{})
}
