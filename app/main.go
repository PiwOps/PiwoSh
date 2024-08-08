// main.go
package main

import (
	"smuggr.xyz/piwosh/common/configurator"
	"smuggr.xyz/piwosh/common/logger"
	"smuggr.xyz/piwosh/core/botter"
	"smuggr.xyz/piwosh/core/datastorer"
)

var Logger = logger.DefaultLogger

func cleanup() {
	Logger.Print("\n")
	Logger.Log(logger.MsgCleaningUp)
	
	botter.Cleanup()
	datastorer.Cleanup()
}

func main() {
	configurator.Initialize()
	logger.Initialize()
	datastorer.Initialize()

	sigChan := botter.Initialize()
	<-sigChan

	cleanup()
}
