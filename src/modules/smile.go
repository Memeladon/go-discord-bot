package modules

import "github.com/bwmarrin/discordgo"

func SmileHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "))" {
		s.ChannelMessageSend(m.ChannelID, "))")
	}
}
