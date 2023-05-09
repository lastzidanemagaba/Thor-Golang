package zidane

import (
	"fmt"
	"time"

	"zidane/common"
)

func DoSomething() {
	fmt.Println("Doing something in Zidane...")

	now := time.Now().In(common.Loc)
	formatting := now.Format("2006-01-02 15:04:05")
	fmt.Println(formatting)
}
