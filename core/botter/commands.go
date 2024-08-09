// botter/commands.go
package botter

import (
	"smuggr.xyz/piwosh/common/logger"

	"github.com/bwmarrin/discordgo"
)

var Commands = make(map[string]*Command)

func PingCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Pong!",
			Flags:   discordgo.MessageFlags(discordgo.MessageFlagsEphemeral),
		},
	})

	if err != nil {
		Logger.Log(logger.ErrRespondingToInteraction.Format(err))
	}
}

func registerCommand(cmd *Command, err error) {
	if err != nil {
		Logger.Log(logger.ErrRegisteringResource.Format(err, logger.ResourceCommand))
		return
	}

	Commands[cmd.DiscordCommand.Name] = cmd

	if cmd.SetupFunc != nil {
		Logger.Debug("Setting up command: %s", cmd.DiscordCommand.Name)
		cmd.SetupFunc(Session)
	}

	Logger.Log(logger.MsgRegisteredResource.Format(cmd.DiscordCommand.Name, logger.ResourceCommand))
}

func setupCommands() {
	registerCommand(NewCommand(&discordgo.ApplicationCommand{
		Name:        "ping",
		Description: "Replies with pong!",
	}, PingCommand, nil))

	registerCommand(NewCommand(&discordgo.ApplicationCommand{
		Name:        "createdynamicvoicechannel",
		Description: "Creates a dynamic voice channel",
	}, CreateDVCCommand, nil))
}