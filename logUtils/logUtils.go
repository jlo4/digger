package logUtils

import (
	"fmt"
	"log"
)

func LogFatalError(err error) bool {
	if err != nil {
		fmt.Println("ERROR")
		log.Fatal(err)
		return true
	}

	return false
}
