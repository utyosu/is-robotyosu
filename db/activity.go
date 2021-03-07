package db

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type Activity struct {
	gorm.Model
	DiscordUserId int64
	Kind          uint32
}

type ActivityKind uint32

const (
	ActivityKindBattlePower         ActivityKind = 1
	ActivityKindFoodPorn            ActivityKind = 2
	ActivityKindFortune             ActivityKind = 3
	ActivityKindInsiderGame         ActivityKind = 4
	ActivityKindInteractionCreate   ActivityKind = 5
	ActivityKindInteractionDestroy  ActivityKind = 6
	ActivityKindInteractionResponse ActivityKind = 7
	ActivityKindInteractionList     ActivityKind = 8
	ActivityKindLuckyColor          ActivityKind = 9
	ActivityKindNickname            ActivityKind = 10
	ActivityKindTalk                ActivityKind = 11
	ActivityKindWeapon              ActivityKind = 12
	ActivityKindWeather             ActivityKind = 13
	ActivityKindLottery             ActivityKind = 14
)

func InsertActivity(discordUserlIdStr string, kind ActivityKind) (*Activity, error) {
	discordUserlId, err := strconv.ParseInt(discordUserlIdStr, 10, 64)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	activity := Activity{
		DiscordUserId: discordUserlId,
		Kind:          uint32(kind),
	}
	if err := dbs.Create(&activity).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return &activity, nil
}

func FetchTodayActivities(discordUserlIdStr string, kind ActivityKind) ([]*Activity, error) {
	now := time.Now()
	todayStartTime := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)

	activities := []*Activity{}
	err := dbs.Find(&activities, "discord_user_id = ? and kind = ? and created_at >= ?", discordUserlIdStr, kind, todayStartTime).Error
	return activities, errors.WithStack(err)
}
