// botter/handlers.go
package botter

import (
    "fmt"
    "os"

    "smuggr.xyz/piwosh/common/logger"

    "github.com/bwmarrin/discordgo"
)

var DynamicChannels = make(map[string]bool)

func isDynamicChannelCreatorChannel(channelID string) bool {
    return channelID == DynamicChannelsCreatorChannelID
}

func isDynamicChannel(channelID string) bool {
    _, exists := DynamicChannels[channelID]
    return exists
}

func createUserVoiceChannel(s *discordgo.Session, userID, categoryID string) (string, error) {
    user, err := s.User(userID)
    if err != nil {
        Logger.Log(logger.ErrFetchingResource.Format(err, logger.ResourceUser))
        return "", err
    }

    guildID := os.Getenv("DISCORD_GUILD_ID")
    channelName := fmt.Sprintf("%s's Channel", user.Username)

    parentID := ""
    if categoryID != "" {
        if parentChannel, err := s.Channel(categoryID); err == nil && parentChannel.Type == discordgo.ChannelTypeGuildCategory {
            parentID = categoryID
        }
    }

    newChannel, err := s.GuildChannelCreateComplex(guildID, discordgo.GuildChannelCreateData{
        Name:     channelName,
        Type:     discordgo.ChannelTypeGuildVoice,
        ParentID: parentID,
    })

    if err != nil {
        Logger.Log(logger.ErrRegisteringResource.Format(err, logger.ResourceChannel))
        return "", err
    }

    DynamicChannels[newChannel.ID] = true

    return newChannel.ID, nil
}

func moveUserToChannel(s *discordgo.Session, guildID, userID, channelID string) {
    err := s.GuildMemberMove(guildID, userID, &channelID)
    if err != nil {
        Logger.Error(err)
    }
}

func checkAndCleanupChannels(s *discordgo.Session) {
    guildID := os.Getenv("DISCORD_GUILD_ID")
    guild, err := s.Guild(guildID)
    if err != nil {
        Logger.Log(logger.ErrFetchingResource.Format(err, logger.ResourceGuild))
        return
    }

    for channelID := range DynamicChannels {
        channel, err := s.Channel(channelID)
        if err != nil {
            Logger.Log(logger.ErrFetchingResource.Format(err, logger.ResourceChannel))
            continue
        }

        if channel.Type != discordgo.ChannelTypeGuildVoice {
            continue
        }

        isEmpty := true
        for _, vs := range guild.VoiceStates {
            if vs.ChannelID == channelID {
                isEmpty = false
                break
            }
        }

        if isEmpty {
            _, err := s.ChannelDelete(channel.ID)
            if err != nil {
                Logger.Log(logger.ErrRemovingResource.Format(err, logger.ResourceChannel))
            } else {
                delete(DynamicChannels, channel.ID)
                Logger.Log(logger.MsgRemovedResource.Format(channel.Name, logger.ResourceChannel))
            }
        }
    }
}

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
    if vsu.BeforeUpdate != nil && isDynamicChannel(vsu.BeforeUpdate.ChannelID) && vsu.ChannelID != vsu.BeforeUpdate.ChannelID {
        checkAndCleanupChannels(s)
    }

    if isDynamicChannelCreatorChannel(vsu.ChannelID) {
        categoryID := getCategoryIDOfChannel(s, vsu.ChannelID)
        newChannelID, err := createUserVoiceChannel(s, vsu.UserID, categoryID)
        if err != nil {
            return
        }

        moveUserToChannel(s, vsu.GuildID, vsu.UserID, newChannelID)
    }
}

func getCategoryIDOfChannel(s *discordgo.Session, channelID string) string {
    channel, err := s.Channel(channelID)
    if err != nil {
        Logger.Log(logger.ErrFetchingResource.Format(err, logger.ResourceChannel))
        return ""
    }

    if channel.ParentID == "" {
        return ""
    }

    parentChannel, err := s.Channel(channel.ParentID)
    if err != nil {
        Logger.Log(logger.ErrFetchingResource.Format(err, logger.ResourceChannel))
        return ""
    }

    if parentChannel.Type == discordgo.ChannelTypeGuildCategory {
        return parentChannel.ID
    } else {
        return ""
    }
}

func setupHandlers() {
    Session.AddHandler(onInteractionCreate)
    Session.AddHandler(onVoiceStateUpdate)
}