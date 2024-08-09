// botter/handlers.go
package botter

import (
	"github.com/bwmarrin/discordgo"
)

func onInteractionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
    if i.Type != discordgo.InteractionApplicationCommand {
        return
    }

    commandName := i.ApplicationCommandData().Name
    if handler, found := Commands[commandName]; found {
        handler.HandlerFunc(s, i)
    }
}

func onVoiceStateUpdate(s *discordgo.Session, vsu *discordgo.VoiceStateUpdate) {
    if vsu.BeforeUpdate != nil && isUserDynamicChannel(vsu.BeforeUpdate.ChannelID) && vsu.ChannelID != vsu.BeforeUpdate.ChannelID {
        checkAndCleanupChannels(s)
    }

    if isDynamicChannelCreatorChannel(vsu.ChannelID) {
        categoryID := getCategoryIDOfChannel(s, vsu.ChannelID)
        newChannelID, err := createUserDynamicChannel(s, vsu.UserID, categoryID)
        if err != nil {
            return
        }

        moveUserToChannel(s, vsu.GuildID, vsu.UserID, newChannelID)
    }
}

func onReady(s *discordgo.Session, r *discordgo.Ready) {
    Logger.Success("bot is running, press ctrl+c to exit")
    setDefaultBotStatus(s)
}

func setupHandlers() {
    Session.AddHandler(onReady)
    Session.AddHandler(onInteractionCreate)
    Session.AddHandler(onVoiceStateUpdate)
}