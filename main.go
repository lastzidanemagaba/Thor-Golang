package main

import (
	"fmt"

	"zidane/common"
	"zidane/zidane"
)

func init() {
	common.InitFunction()
}

func main() {
	fmt.Println("Running main...")

	zidane.DoSomething()
}
