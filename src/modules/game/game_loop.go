package game

import (
	"fmt"
	"math/rand"

	"github.com/bwmarrin/discordgo"
)

func pvpSimulation(player1, player2 *Player, s *discordgo.Session, m *discordgo.MessageCreate) {
	minDamage, maxDamage := 2, 5
	minHealth, maxHealth, minManaCost, maxManaCost := 6, 7, 1, 2
	minMana, maxMana := 4, 6

	gameMove := 0
	// Цикл игры, пока кто-то не победит
	for gameMove != 6 {
		gameMove++

		messagegameMove := fmt.Sprintf("\nХод %d:\n", gameMove)

		// Ход первого игрока
		action := rand.Intn(3)
		switch action {
		case 0:
			player1.Attack(player2, minDamage, maxDamage)
		case 1:
			player1.Heal(minHealth, maxHealth, minManaCost, maxManaCost)
		case 2:
			player1.Refill(minMana, maxMana)
		}

		// Проверяем состояние второго игрока
		if !player2.IsAlive() {
			messageWin := fmt.Sprintf("%s победил!\n", player1.Username)
			s.ChannelMessageSend(m.ChannelID, messageWin)
			return
		}

		// Ход второго игрока
		action = rand.Intn(3)
		switch action {
		case 0:
			player2.Attack(player1, minDamage, maxDamage)
		case 1:
			player2.Heal(minHealth, maxMana, minManaCost, maxManaCost)
		case 2:
			player2.Refill(minMana, maxMana)
		}

		// Проверяем состояние первого игрока
		if !player1.IsAlive() {
			messageWin := fmt.Sprintf("%s победил!\n", player2.Username)
			s.ChannelMessageSend(m.ChannelID, messageWin)
			return
		}

		// Выводим текущее состояние
		messagePlayer1 := fmt.Sprintf("Здоровье %s: %d (Мана: %d)\n",
			player1.Username, player1.Health, player1.Mana)
		messagePlayer2 := fmt.Sprintf("Здоровье %s: %d (Мана: %d)\n",
			player2.Username, player2.Health, player2.Mana)
		messageTotal := messagegameMove + messagePlayer1 + messagePlayer2

		s.ChannelMessageSend(m.ChannelID, messageTotal)

	}
}
