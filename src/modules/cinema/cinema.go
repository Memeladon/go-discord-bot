package cinema

import (
	"fmt"
	"go-bot/src/constants"
	"go-bot/src/helpers"
	"net/url"
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func Handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	args := helpers.ParseCommand(m)

	//TODO: переделать передачу ивентов в хэндлеры
	if len(args) < 1 || args[0] != "cinema" {
		return
	}

	// TODO: убрать костыль, без него пока может упасть если юзер не допишет команду
	// ожидается add TITLE DESCRIPTION или ссылка на кинопоиск
	if len(args) < 3 {
		s.ChannelMessageSend(m.ChannelID, "не понял")
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
		// https://www.kinopoisk.ru/film/1319995/

		if u, err := url.ParseRequestURI(args[2]); err == nil && helpers.IsWhitelistedHost(u) {
			movie := KinopoiskMovie{}

			re, _ := regexp.Compile(`[0-9]+`)
			movieIndex := re.FindString(u.EscapedPath())

			err := movie.Init(movieIndex)
			if err != nil {
				s.ChannelMessageSend(m.ChannelID, err.Error())
				return
			}

			var genres []string
			for _, genre := range movie.Genres {
				genres = append(genres, genre.Name)
			}

			embeds = append(embeds, &discordgo.MessageEmbed{
				Image: &discordgo.MessageEmbedImage{
					URL:      movie.Poster.Url,
					ProxyURL: movie.Poster.PreviewUrl,
				},
			})

			hyperLinkName := fmt.Sprintf("[%s](%s)", movie.Name, u)

			s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{
				Content: composeMoviePost(
					hyperLinkName,
					movie.Description,
					fmt.Sprint(movie.Year),
					genres,
					fmt.Sprint(movie.MovieLength),
					m.Author.Mention(),
				),
				Embeds: embeds,
			})

		} else if len(args) > 3 {
			//TODO: передавать дату и время
			s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{
				Content: composeMoviePost(
					args[2],
					args[3],
					"1970",
					[]string{""},
					"0",
					m.Author.Mention(),
				),
				Embeds: embeds,
			})
		} else {
			s.ChannelMessageSend(m.ChannelID, helpers.GetRandomElement(constants.DontUnderstandAnswers[:]))
		}
	}
}

func composeMoviePost(title, description, year string, genres []string, time, mention string) string {
	return fmt.Sprintf(
		"**%s**\n*%s*\n\nгод: %s\nжанры: %s\nпродолжительность: %s мин\nпредложил: %s",
		strings.ToUpper(title),
		description,
		year,
		fmt.Sprint(genres),
		time,
		mention,
	)
}
