package cinema

import (
	"fmt"
	"go-bot/src/constants"
	"go-bot/src/helpers"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func Handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	args := helpers.ParseCommand(m)

	if args[0] != "cinema" {
		return
	}

	// TODO: убрать костыль, без него пока может упасть если юзер не допишет команду
	// ожидается add TITLE DESCRIPTION
	if len(args) < 4 {
		s.ChannelMessageSend(m.ChannelID, "не понял")
		return
	}

	cinemaRole, err := helpers.FindRoleInGuildByName(constants.CinemaRole, m.GuildID, s)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, err.Error())
		return
	}

	var embeds []*discordgo.MessageEmbed
	for _, v := range m.Attachments {
		embeds = append(embeds, &discordgo.MessageEmbed{
			Image: &discordgo.MessageEmbedImage{
				URL:      v.URL,
				ProxyURL: v.ProxyURL,
			},
		})
	}

	switch args[1] {
	case "add":
		s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{
			Content: composeMoviePost(args[2], args[3], cinemaRole.Mention()),
			Embeds:  embeds,
		})
	}
}

func composeMoviePost(title, description, cinemaRoleMention string) string {
	return fmt.Sprintf(`
	> **%s**
	> *%s*
	> %s смотрим?
	`, strings.ToUpper(title), description, cinemaRoleMention)
}
