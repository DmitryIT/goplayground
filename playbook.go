package main

import (
	"errors"
	"fmt"
	"log"
)

const GITHUB_TOCKEN string = "hHGhsdfj6adshHh83j4"
const VAULT_PASSWORD string = "sdfsdfjnsdkfjb/nbsdfnb/mnbsdf"

type Counter struct {
	total int
}

func (c Counter) String() {
	fmt.Println("Counter = ", c.total)
}

func (c Counter) SetCounter(val int) error {
	if val < 0 {
		err := errors.New("total is negative")
		return err
	}
	c.total = val
	return nil
}

func (c *Counter) SetCounterPrt(val int) error {
	if val < 0 {
		err := errors.New("total is negative")
		return err
	}
	c.total = val
	return nil
}

func main() {
	var cnt Counter
	cnt.String()

	if err := cnt.SetCounter(1); err != nil {
		log.Fatalln(err)
	}
	cnt.String()

	if err := cnt.SetCounterPrt(10); err != nil {
		log.Fatalln(err)
	}
	cnt.String()
}
