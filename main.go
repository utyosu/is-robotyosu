package main

import (
	"github.com/utyosu/rfe/app"
	"github.com/utyosu/rfe/db"
	"github.com/utyosu/rfe/env"
	"github.com/utyosu/robotyosu-go/slack"
	"math/rand"
	"time"
)

func init() {
	if loc, err := time.LoadLocation("Asia/Tokyo"); err == nil {
		time.Local = loc
	}
	rand.Seed(time.Now().UnixNano())
	slack.Init(&slack.SlackConfig{
		WarningChannel: env.SlackChannelWarning,
		AlertChannel:   env.SlackChannelAlert,
		Token:          env.SlackToken,
		Title:          env.SlackTitle,
	})
}

func main() {
	defer app.NotifySlackWhenPanic("main")
	db.ConnectDb()
	app.Start()
	return
}
