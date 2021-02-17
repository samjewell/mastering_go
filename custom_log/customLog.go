package main

import (
	"fmt"
	"log"
	"os"
)

// LOGFILE is where the logs are going to land
var LOGFILE = "/tmp/mGo.log"

func main() {
	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	// LstdFlags = Ldate | Ltime
	// initial values for the standard logger
	iLog := log.New(f, "customLog ", log.LstdFlags)
	iLog.SetFlags(log.LstdFlags)
	iLog.Println("Hello there!")
	iLog.Println("Another log entry!")
}
