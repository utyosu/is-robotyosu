package app

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
	"github.com/utyosu/rfe/db"
	"math/rand"
)

func isLotteryExecute(m *discordgo.MessageCreate) bool {
	return isContainKeywords(m.Content, lotteryKeywords)
}

func actionLottery(m *discordgo.MessageCreate) {
	activities, err := db.FetchTodayActivities(m.Author.ID, db.ActivityKindLottery)
	if err != nil {
		sendMessage(m.ChannelID, messageError)
		return
	} else if len(activities) >= 3 {
		sendMessage(m.ChannelID, "くじは1日に3回までしか引けません。また明日チャレンジしてね！")
		return
	}

	var msg string
	r := rand.Intn(6096454)
	if r == 0 {
		msg = "%vさん、1等が当たりました！当選金額は2億円です！"
	} else if r <= 6 {
		msg = "%vさん、2等が当たりました！当選金額は1000万円です！"
	} else if r <= 222 {
		msg = "%vさん、3等が当たりました！30万円です！"
	} else if r <= 10212 {
		msg = "%vさん、4等が当たりました！当選金額は6800円です！"
	} else if r <= 165612 {
		msg = "%vさん、5等が当たりました！当選金額は1000円です！"
	} else {
		msg = "%vさん、はずれです(´・ω・`)ｼｮﾎﾞﾎﾞｰﾝ"
	}
	sendMessage(
		m.ChannelID,
		fmt.Sprintf(msg, getName(m)),
	)

	if _, err := db.InsertActivity(m.Author.ID, db.ActivityKindLottery); err != nil {
		slackWarning.Post(errors.WithStack(err))
	}
}

var (
	lotteryKeywords = []string{
		"くじ",
	}
)
