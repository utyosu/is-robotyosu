package app

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func isHelpExecute(m *discordgo.MessageCreate) bool {
	return isContainKeywords(m.Content, helpKeywords)
}

func actionHelp(m *discordgo.MessageCreate) {
	var msg string
	for _, v := range helpCommands {
		if len(v.keywords) > 0 {
			msg += fmt.Sprintf("**`%v`**→ %v\n", v.keywords[0], v.description)
		}
	}
	sendMessage(m.ChannelID, msg)
}

type helpCommand struct {
	keywords    []string
	description string
}

var (
	helpKeywords = []string{
		"help",
		"使い方",
	}
	helpCommands = []helpCommand{
		{weaponKeywords, "スプラトゥーン2のオススメブキを表示"},
		{foodPornKeywords, "ごちそうを表示"},
		{battlePowerKeyword, "戦闘力を測定"},
		{lotteryKeywords, "宝くじを引く"},
	}
)
