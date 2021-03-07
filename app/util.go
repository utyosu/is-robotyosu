package app

import (
	"github.com/bwmarrin/discordgo"
	"strings"
	"time"
)

func isContainKeywords(m string, keywords []string) bool {
	for _, k := range keywords {
		if strings.Contains(m, k) {
			return true
		}
	}
	return false
}

func getName(m *discordgo.MessageCreate) string {
	if m.Member != nil && m.Member.Nick != "" {
		return m.Member.Nick
	}
	return m.Author.Username
}

func NotifySlackWhenPanic(p ...interface{}) {
	if err := recover(); err != nil {
		slackAlert.Post(p...)
	}
}

func doFuncSchedule(f func(), interval time.Duration) *time.Ticker {
	ticker := time.NewTicker(interval)
	go func() {
		for range ticker.C {
			f()
		}
	}()
	return ticker
}
