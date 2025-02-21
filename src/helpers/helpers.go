package helpers

import (
	"go-bot/src/constants"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func ParseCommand(m *discordgo.MessageCreate) []string {
	withiout := m.Content[len(constants.CommandPrefix):]
	return strings.Split(withiout, " ")

}
