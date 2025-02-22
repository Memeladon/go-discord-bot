package smile

import (
	"github.com/bwmarrin/discordgo"
)

func Handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content == "))" {
		s.ChannelMessageSend(m.ChannelID, "))")
	}
}
