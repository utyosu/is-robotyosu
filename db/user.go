package db

import (
	basic_errors "errors"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"strconv"
)

type User struct {
	gorm.Model
	DiscordUserId int64
	Name          string
}

func FindUser(discordUserIdStr string) (*User, error) {
	user := User{}
	if err := dbs.Take(&user, "discord_user_id=?", discordUserIdStr).Error; err != nil {
		if basic_errors.Is(err, gorm.ErrRecordNotFound) {
			return &user, nil
		}
		return nil, errors.WithStack(err)
	}
	return &user, nil
}

func FindOrCreateUser(discordUserIdStr string, name string) (*User, error) {
	user, err := FindUser(discordUserIdStr)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if user.ID == 0 {
		discordUserId, _ := strconv.ParseInt(discordUserIdStr, 10, 64)
		user.DiscordUserId = discordUserId
		user.Name = name
		if err := dbs.Create(&user).Error; err != nil {
			return nil, errors.WithStack(err)
		}
	} else if user.Name != name {
		user.Name = name
		if err := dbs.Save(&user).Error; err != nil {
			return nil, errors.WithStack(err)
		}
	}
	return user, nil
}
