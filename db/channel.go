package db

import (
	basic_errors "errors"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Channel struct {
	gorm.Model
	DiscordChannelId int64
	DiscordGuildId   int64
	Name             string
}

func FindChannel(discordChannelId int64) (*Channel, error) {
	channel := Channel{}
	if err := dbs.Take(&channel, "discord_channel_id=?", discordChannelId).Error; err != nil {
		if basic_errors.Is(err, gorm.ErrRecordNotFound) {
			return &channel, nil
		}
		return nil, errors.WithStack(err)
	}
	return &channel, nil
}

func FindOrCreateChannel(discordChannelId, discordGuildId int64, name string) (*Channel, error) {
	channel, err := FindChannel(discordChannelId)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if channel.ID == 0 {
		channel.DiscordChannelId = discordChannelId
		channel.DiscordGuildId = discordGuildId
		channel.Name = name
		if err := dbs.Create(&channel).Error; err != nil {
			return nil, errors.WithStack(err)
		}
	} else if channel.Name != name {
		channel.Name = name
		if err := dbs.Save(&channel).Error; err != nil {
			return nil, errors.WithStack(err)
		}
	}
	return channel, nil
}
