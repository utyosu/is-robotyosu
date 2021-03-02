package app

import (
	"github.com/bwmarrin/discordgo"
	"strings"
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
