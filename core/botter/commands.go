// botter/commands.go
package botter

import (
	"fmt"

	"smuggr.xyz/piwosh/common/logger"

	"github.com/bwmarrin/discordgo"
)

var Commands = make(map[string]*Command)
var DynamicChannelsCreatorChannelID string

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

func CreateDynamicVoiceChannelCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	guildID := i.GuildID
	channelID := i.ChannelID
	categoryID := getCategoryIDOfChannel(s, channelID)

	channel, err := s.GuildChannelCreateComplex(guildID, discordgo.GuildChannelCreateData{
		Name:     "Dynamic Voice Channel",
		Type:     discordgo.ChannelTypeGuildVoice,
		ParentID: categoryID,
	})

	if err != nil {
		Logger.Log(logger.ErrRegisteringResource.Format(err, logger.ResourceChannel))
		return
	}

	DynamicChannelsCreatorChannelID = channel.ID

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("Created dynamic voice channel: <#%s>", channel.ID),
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
	Logger.Log(logger.MsgRegisteredResource.Format(cmd.DiscordCommand.Name, logger.ResourceCommand))
}

func setupCommands() {
	registerCommand(NewCommand(&discordgo.ApplicationCommand{
		Name:        "ping",
		Description: "Replies with pong!",
	}, PingCommand))

	registerCommand(NewCommand(&discordgo.ApplicationCommand{
		Name:        "createdynamicvoicechannel",
		Description: "Creates a dynamic voice channel",
	}, CreateDynamicVoiceChannelCommand))
}