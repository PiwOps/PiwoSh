package botter

import (
	"smuggr.xyz/piwosh/common/logger"

	"github.com/bwmarrin/discordgo"
)

func isValidChannel(s *discordgo.Session, channelID string) bool {
    _, err := s.Channel(channelID)
    return err == nil
}

func moveUserToChannel(s *discordgo.Session, guildID, userID, channelID string) {
    err := s.GuildMemberMove(guildID, userID, &channelID)
    if err != nil {
        Logger.Error(err)
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

func setDefaultBotStatus(s *discordgo.Session) error {
	return s.UpdateCustomStatus(Config.DefaultStatus)
}