package db

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type UserStatus struct {
	gorm.Model
	DiscordUserId    int64
	DiscordChannelId int64
	IntervalTime     int64
}

func InsertUserStatus(discordUserId, discordChannelId, intervalTime int64) error {
	userStatus := UserStatus{
		DiscordUserId:    discordUserId,
		DiscordChannelId: discordChannelId,
		IntervalTime:     intervalTime,
	}
	if err := dbs.Save(&userStatus).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}
