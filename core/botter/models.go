// botter/models.go
package botter

import (
	"os"

	"github.com/bwmarrin/discordgo"
)

type BotState struct {

}

type Command struct {
	DiscordCommand *discordgo.ApplicationCommand
	HandlerFunc    func(s *discordgo.Session, i *discordgo.InteractionCreate)
}

func NewCommand(discordCmd *discordgo.ApplicationCommand, handlerFunc func(s *discordgo.Session, i *discordgo.InteractionCreate)) (*Command, error) {
	appCmd, err := Session.ApplicationCommandCreate(os.Getenv("DISCORD_APP_ID"), os.Getenv("DISCORD_GUILD_ID"), discordCmd)
	if err != nil {
		return nil, err
	}

	return &Command{
		DiscordCommand: appCmd,
		HandlerFunc: 	handlerFunc,
	}, nil
}