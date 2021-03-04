package main

import (
	"github.com/utyosu/rfe/app"
	"github.com/utyosu/rfe/db"
	"math/rand"
	"time"
)

func init() {
	if loc, err := time.LoadLocation("Asia/Tokyo"); err == nil {
		time.Local = loc
	}
	rand.Seed(time.Now().UnixNano())
}

func main() {
	defer app.NotifySlackWhenPanic("main")
	db.ConnectDb()
	app.Start()
	return
}
