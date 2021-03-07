package app

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func isVersionExecute(m *discordgo.MessageCreate) bool {
	return m.Content == ".rfe version"
}

func actionVersion(m *discordgo.MessageCreate) {
	sendMessage(m.ChannelID, fmt.Sprintf(
		"CommitHash: %v\nBuildDatetime: %v",
		commitHash,
		buildDatetime,
	))
}
