package common

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var Loc *time.Location
var Now time.Time
var Tz string
var Mendua string
var pushtoproduction = false // Set to false to use .env file in development
var file *os.File
var err error

func InitFunction() {

	if pushtoproduction == false {
		file, err = os.Open(".env")
		if err != nil {
			log.Fatalf("Error opening .env file: %v", err)
		}
		defer file.Close()
		// Read the file line by line
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			// Split each line into key-value pairs
			pair := strings.SplitN(scanner.Text(), "=", 2)

			// Set the key-value pair as an environment variable
			if err := os.Setenv(pair[0], pair[1]); err != nil {
				log.Fatalf("Error setting environment variable: %v", err)
			}
		}

		if err := scanner.Err(); err != nil {
			log.Fatalf("Error reading .env file: %v", err)
		}
	}

	Tz = os.Getenv("TIMEZONE")
	Mendua = os.Getenv("DATA")
	Loc, err = time.LoadLocation(Tz)
	if err != nil {
		fmt.Println(err)
		return
	}

	Now = time.Now().In(Loc)
}
