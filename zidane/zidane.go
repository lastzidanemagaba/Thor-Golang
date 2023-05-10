package zidane

import (
	"fmt"
	"zidane/common"
)

func DoSomething() {
	fmt.Println("Doing something in Zidane...")

	fmt.Println("Current time in Jakarta:", common.Now.Format("2006-01-02 15:04:05"))

	fmt.Println("TZnya ", common.Tz)
}
