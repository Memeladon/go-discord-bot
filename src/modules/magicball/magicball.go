package magicball

import (
	"go-bot/src/helpers"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var possibleAnswers = [...]string{
	"определённо",
	"так точно",
	"без сомнений да",
	"да, это так",
	"ты можешь довериться, что это правда",
	"насколько я вижу - да",
	"скорее всего",
	"походу так",
	"да",
	"знаки говорят, что да",
	"звёзды не сошлись, ответ неясен",
	"давай потом, ладно?",
	"я лучше промолчу",
	"не могу подсказать сейчас",
	"лучше подумай и спроси ещё раз",
	"лол, нет",
	"мой ответ - нет",
	"мои источники говорят нет",
	"да вроде нет",
	"очень сомнительно",
}

func Handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// TODO: хэлпер для парсинга команд
	if strings.HasPrefix(m.Content, "))question") {
		s.ChannelMessageSend(m.ChannelID, helpers.GetRandomElement(possibleAnswers[:]))
	}
}
