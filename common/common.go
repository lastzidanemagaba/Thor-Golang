package common

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

var Loc *time.Location
var Now time.Time
var Tz string

func InitFunction() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	}
	Tz = os.Getenv("TIMEZONE")
	Loc, err = time.LoadLocation(Tz)
	if err != nil {
		fmt.Println(err)
		return
	}

	Now = time.Now().In(Loc)
}
