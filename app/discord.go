package app

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
	"github.com/utyosu/rfe/db"
	"github.com/utyosu/rfe/env"
	"github.com/utyosu/robotyosu-go/slack"
	"log"
)

var (
	discordSession *discordgo.Session
	stopBot        = make(chan bool)
	slackAlert     *slack.Config
	slackWarning   *slack.Config
	commitHash     string
	buildDatetime  string
)

func init() {
	slackAlert = &slack.Config{
		Channel: env.SlackChannelAlert,
		Token:   env.SlackToken,
		Title:   env.SlackTitle,
	}
	slackWarning = &slack.Config{
		Channel: env.SlackChannelWarning,
		Token:   env.SlackToken,
		Title:   env.SlackTitle,
	}
}

func Start() {
	var err error
	discordSession, err = discordgo.New(fmt.Sprintf("Bot %s", env.DiscordBotToken))
	if err != nil {
		panic(err)
	}

	discordSession.AddHandler(onMessageCreate)
	if err = discordSession.Open(); err != nil {
		panic(err)
	}
	defer discordSession.Close()
	log.Println("Listening...")

	doFuncSchedule(recordServerActivities, env.RecordInterval)
	<-stopBot
	return
}

func onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	defer NotifySlackWhenPanic(s, m)

	// 自分のメッセージは処理しない
	if m.Author.ID == env.DiscordBotClientId {
		return
	}

	// 対象のチャンネル以外のメッセージは処理しない
	if m.ChannelID != env.DiscordTargetChannelId {
		return
	}
	log.Printf("\t%v\t%v\t%v\t%v\t%v\n", m.GuildID, m.ChannelID, m.Type, m.Author.Username, m.Content)

	// ユーザー登録
	name := m.Author.Username
	if m.Member != nil && m.Member.Nick != "" {
		name = m.Member.Nick
	}
	if _, err := db.FindOrCreateUser(m.Author.ID, name); err != nil {
		slackWarning.Post(err)
	}

	switch {
	case isVersionExecute(m):
		actionVersion(m)
	case isHelpExecute(m):
		actionHelp(m)
	case isWeaponExecute(m):
		actionWeapon(m)
	case isBattlePowerExecute(m):
		actionBattlePower(m)
	case isFoodPornExecute(m):
		actionFoodPorn(m)
	case isLotteryExecute(m):
		actionLottery(m)
	}
}

func sendMessage(channelID string, msg string) {
	if _, err := discordSession.ChannelMessageSend(channelID, msg); err != nil {
		slackWarning.Post(
			errors.WithStack(err),
			channelID,
			msg,
		)
	}
}
