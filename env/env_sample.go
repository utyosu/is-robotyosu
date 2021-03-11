// +build sample

package env

import (
	"time"
)

const (
	DiscordBotToken        = ""
	DiscordBotClientId     = ""
	DbDriver               = "mysql"
	DbUser                 = ""
	DbPassword             = ""
	DbHost                 = "127.0.0.12"
	DbPort                 = "3306"
	DbName                 = ""
	DbLogLevel             = "warn" // silent, error, warn, info
	DiscordTargetChannelId = ""
	GoogleSearchApiKey     = ""
	GoogleSearchApiCx      = ""
	SlackToken             = ""
	SlackChannelWarning    = "#channel-name"
	SlackChannelAlert      = "#channel-name"
	SlackTitle             = ""
	RecordGuildId          = ""
	RecordInterval         = time.Second * 300
)
