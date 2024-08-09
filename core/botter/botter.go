// botter.go
package botter

import (
	"os"
	"os/signal"
	"syscall"

	"smuggr.xyz/piwosh/common/logger"
	"smuggr.xyz/piwosh/common/configurator"

	"github.com/bwmarrin/discordgo"
)

var Logger = logger.NewCustomLogger("bott")
var Config *configurator.BotConfig
var Session *discordgo.Session

func Initialize() chan os.Signal {
	Logger.Info(logger.MsgInitializing)
	Config = &configurator.Config.Bot

	s, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		Logger.Fatalf("invalid session parameters: %v", err)
		return nil
	}
	Session = s

	setupHandlers()

	err = Session.Open()
	if err != nil {
		Logger.Fatalf("error opening discord session: %v", err)
		return nil
	}

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	setupCommands()

	return signalCh
}

func Cleanup() {
	Logger.Log(logger.MsgCleaningUp)
	
	checkAndCleanupChannels(Session)
	Session.Close()
}
