package common

import (
	"os"
	"time"
)

var Loc *time.Location

func InitFunction() {
	var err error
	tz := os.Getenv("TIMEZONE")
	Loc, err = time.LoadLocation(tz)
	if err != nil {
		panic(err)
	}
}
