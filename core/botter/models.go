// botter/models.go
package botter

import (
	"os"

	"github.com/bwmarrin/discordgo"
)

type Command struct {
	DiscordCommand *discordgo.ApplicationCommand
	HandlerFunc    func(s *discordgo.Session, i *discordgo.InteractionCreate)
	SetupFunc      func(s *discordgo.Session)
}

func NewCommand(discordCmd *discordgo.ApplicationCommand, handlerFunc func(s *discordgo.Session, i *discordgo.InteractionCreate), setupFunc func(s *discordgo.Session)) (*Command, error) {
	appCmd, err := Session.ApplicationCommandCreate(os.Getenv("DISCORD_APP_ID"), os.Getenv("DISCORD_GUILD_ID"), discordCmd)
	if err != nil {
		return nil, err
	}

	return &Command{
		DiscordCommand: appCmd,
		HandlerFunc: 	handlerFunc,
		SetupFunc:      setupFunc,
	}, nil
}