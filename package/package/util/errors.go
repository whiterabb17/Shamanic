package util

import (
	"log"
)

// Calm consumes top-level errors.
// Defer this if you want to be absolutely certain that local errors should be ignored.
func Calm() {
	if r := recover(); r != nil {
		log.Println("Recovered from", r)
	}
}

func Gate(recp string) bool {
	if recp == "all" || recp == ID {
		return true
	}
	return false
}

// Handle will panic when given a valid error and print some debug info.
func Handle(err error, str ...interface{}) {
	if err == nil {
		return
	}
	if len(str) > 0 {
		log.Println(str[0].(string)+",", err)
		//log.Panicln(str[0].(string)+",", err)
	} else {
		log.Println(err)
		//		log.Panicln(err)
	}
}
