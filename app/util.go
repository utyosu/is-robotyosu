package app

import (
	"fmt"
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

func NotifySlackWhenPanic(info string) {
	if err := recover(); err != nil {
		postSlackAlert(fmt.Sprintf("panic: %v\ninfo: %v", err, info))
	}
}

func messageInformation(s *discordgo.Session, m *discordgo.MessageCreate) string {
	return fmt.Sprintf(
		"session: %v\nmessage: %v",
		s,
		m,
	)
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
