package app

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/utyosu/rfe/env"
	"log"
	"time"
)

var (
	discordSession *discordgo.Session
	stopBot        = make(chan bool)
)

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

	doFuncSchedule(recordServerActivities, time.Second*env.RecordIntervalTime)
	<-stopBot
	return
}

func onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	defer NotifySlackWhenPanic(messageInformation(s, m))

	// 自分のメッセージは処理しない
	if m.Author.ID == env.DiscordBotClientId {
		return
	}

	// 対象のチャンネル以外のメッセージは処理しない
	if m.ChannelID != env.DiscordTargetChannelId {
		return
	}
	log.Printf("\t%v\t%v\t%v\t%v\t%v\n", m.GuildID, m.ChannelID, m.Type, m.Author.Username, m.Content)

	switch {
	case isHelpExecute(m):
		actionHelp(m)
	case isWeaponExecute(m):
		actionWeapon(m)
	case isBattlePowerExecute(m):
		actionBattlePower(m)
	case isFoodPornExecute(m):
		actionFoodPorn(m)
	}
}

func sendMessage(channelID string, msg string) {
	if _, err := discordSession.ChannelMessageSend(channelID, msg); err != nil {
		postSlackWarning(fmt.Sprintf("Error sending message: %v", err))
	}
}
