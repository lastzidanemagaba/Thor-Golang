package main

import (
	"fmt"
	"time"

	"zidane/common"
	"zidane/zidane"
)

func init() {
	common.InitFunction()
}

func main() {
	fmt.Println("Running main...")
	loca := time.Local
	Nowa := time.Now().In(loca)
	formatting := Nowa.Format("2006-01-02 15:04:05")
	fmt.Println("Local timezone:", formatting)
	zidane.DoSomething()
}
