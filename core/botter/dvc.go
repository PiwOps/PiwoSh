package botter

import (
	"fmt"
	"os"

	"smuggr.xyz/piwosh/common/logger"
	"smuggr.xyz/piwosh/core/datastorer"

	"github.com/bwmarrin/discordgo"
)

var UserDynamicChannels = make(map[string]bool)

func isDynamicChannelCreatorChannel(channelID string) bool {
	var dvccchannelId string
	datastorer.GetCmdConfigValue("DVCCChannelID", &dvccchannelId)
	
    return channelID == dvccchannelId
}

func isUserDynamicChannel(channelID string) bool {
    _, exists := UserDynamicChannels[channelID]
    return exists
}

func createUserDynamicChannel(s *discordgo.Session, userID, categoryID string) (string, error) {
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

    UserDynamicChannels[newChannel.ID] = true

    return newChannel.ID, nil
}

func checkAndCleanupChannels(s *discordgo.Session) {
    guildID := os.Getenv("DISCORD_GUILD_ID")
    guild, err := s.Guild(guildID)
    if err != nil {
        Logger.Log(logger.ErrFetchingResource.Format(err, logger.ResourceGuild))
        return
    }

    for channelID := range UserDynamicChannels {
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
                delete(UserDynamicChannels, channel.ID)
                Logger.Log(logger.MsgRemovedResource.Format(channel.Name, logger.ResourceChannel))
            }
        }
    }
}

func CreateDVCCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	guildID := i.GuildID
	channelID := i.ChannelID
	categoryID := getCategoryIDOfChannel(s, channelID)

    var dvccchannelId string
    datastorer.GetCmdConfigValue("DVCCChannelID", &dvccchannelId)
    if isValidChannel(s, dvccchannelId) {
        err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
            Type: discordgo.InteractionResponseChannelMessageWithSource,
            Data: &discordgo.InteractionResponseData{
                Content: fmt.Sprintf("A dynamic voice channel already exists: <#%s>", dvccchannelId),
                Flags:   discordgo.MessageFlags(discordgo.MessageFlagsEphemeral),
            },
        })

        if err != nil {
            Logger.Log(logger.ErrRespondingToInteraction.Format(err))
        }

        return
    }

	channel, err := s.GuildChannelCreateComplex(guildID, discordgo.GuildChannelCreateData{
		Name:     "Dynamic Voice Channel",
		Type:     discordgo.ChannelTypeGuildVoice,
		ParentID: categoryID,
	})

	if err != nil {
		Logger.Log(logger.ErrRegisteringResource.Format(err, logger.ResourceChannel))
		return
	}

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

	datastorer.SetCmdConfigValue("DVCCChannelID", channel.ID)
}