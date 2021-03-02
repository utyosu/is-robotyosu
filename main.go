package main

import (
	"github.com/utyosu/rfe/app"
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
	app.Start()
	return
}
