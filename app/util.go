package app

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"runtime"
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
		stackTrace := []string{}
		for depth := 0; ; depth++ {
			_, file, line, ok := runtime.Caller(depth)
			if !ok {
				break
			}
			stackTrace = append(stackTrace, fmt.Sprintf("%v: %v:%v", depth, file, line))
		}
		p = append(p[:2], p[0:]...)
		p[0] = err
		p[1] = stackTrace
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
