package main

import "log"

func check(e error, msg string) {
	if e != nil {
		log.Fatalln(msg, e)
		panic(e)
	}
}

func checkBool(e bool, msg string) {
	if e == true {
		log.Fatalln(msg, e)
		panic(e)
	}
}

func defaultTo(boolErr bool, givenValue string, defaultValue string) string {
	if boolErr == true {
		return defaultValue
	}
	return givenValue
}
