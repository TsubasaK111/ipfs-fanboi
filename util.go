package main

import "log"

func check(e error, msg string) {
	if e != nil {
		log.Fatalln(msg, e)
		panic(e)
	}
}
