package main

import (
	"top.moma.go.simple/mlogger"
	"top.moma.go.simple/webserver"
)

func main() {
	mlogger.InfoLogger.Println("Application Start...")
	webserver.HandleRequest()
}
