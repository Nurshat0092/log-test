package main

import log "github.com/sirupsen/logrus"

func main() {
	log.Println("This is a log line")
	log.Println("Another log line")
	log.Fatal("This is really bad, exiting...")
}
