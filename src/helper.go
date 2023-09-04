package main

import "log"

func handleError(e error, note string) {
	log.Println(note)
	log.Fatalln(e)
}
