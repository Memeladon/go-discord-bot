package game

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func SimHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.HasPrefix(m.Content, "))") {

		// Создаём игроков с нужными параметрами
		player1 := NewPlayer("@memeladon", 30, 10)
		player2 := NewPlayer("@lardira", 30, 10)

		// Запускаем симуляцию
		winner := pvp_simulation(player1, player2)
		fmt.Printf("Победитель: %s\n", winner)
	}
}

func pvp_simulation(player1, player2 *Player) string {
	minDamage, maxDamage := 2, 5
	minHealth, maxHealth, minManaCost, maxManaCost := 6, 7, 1, 2
	minMana, maxMana := 4, 6

	gameMove := 0
	// Цикл игры, пока кто-то не победит
	for gameMove != 6 {
		gameMove++
		fmt.Printf("\nХод %d:\n", gameMove)

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
			fmt.Printf("%s победил!\n", player1.Username)
			return player1.Username
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
			fmt.Printf("%s победил!\n", player2.Username)
			return player2.Username
		}

		// Выводим текущее состояние
		fmt.Printf("Здоровье %s: %d (Мана: %d)\n",
			player1.Username, player1.Health, player1.Mana)
		fmt.Printf("Здоровье %s: %d (Мана: %d)\n",
			player2.Username, player2.Health, player2.Mana)
	}
}
