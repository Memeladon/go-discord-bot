package middleware

import (
	"go-bot/src/constants"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type Handler func(*discordgo.Session, *discordgo.MessageCreate)

func CheckCommandMiddleware(handler Handler) func(*discordgo.Session, *discordgo.MessageCreate) {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if !strings.HasPrefix(m.Content, constants.CommandPrefix) {
			return
		}

		handler(s, m)
	}
}

func IgnoreSelfMiddleware(handler Handler) func(*discordgo.Session, *discordgo.MessageCreate) {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		handler(s, m)
	}
}
