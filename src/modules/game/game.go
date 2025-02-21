package game

import (
	"go-bot/src/helpers"

	"github.com/bwmarrin/discordgo"
)

func Handler(s *discordgo.Session, m *discordgo.MessageCreate) {

	args := helpers.ParseCommand(m)

	if args[0] == "game" {

		// Создаём игроков с параметрами
		player1 := NewPlayer("@memeladon", 30, 10)
		player2 := NewPlayer("@lardira", 30, 10)

		// Запускаем симуляцию
		pvpSimulation(player1, player2, s, m)
	}
}
