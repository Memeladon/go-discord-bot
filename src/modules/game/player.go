package game

import (
	"math/rand"
)

type Player struct {
	Username  string
	HealthCap int
	Health    int
	ManaCap   int
	Mana      int
}

func NewPlayer(username string, hp, mp int) *Player {
	return &Player{
		Username:  username,
		HealthCap: hp, // Максимум здоровья
		Health:    hp, // Начальное здоровье
		ManaCap:   mp, // Максимум здоровья
		Mana:      mp, // Начальная мана
	}
}

/* Основные действия игрока */

// Наносит урон вражескому игроку/существу
func (p *Player) Attack(target *Player, minDamage, maxDamage int) {
	damage := rand.Intn(maxDamage-1) + minDamage // Урон от minDamage до maxDamage
	target.Health -= damage

	// fmt.Printf("%s атаковал %s на %d урона!\n",
	// 	p.Username, target.Username, damage)

	if target.Health < 0 {
		target.Health = 0
	}
}

// Восстанавливает ману игроку
func (p *Player) Refill(minMana, maxMana int) {
	restoredMana := rand.Intn(maxMana-minMana+1) + minMana // Мана от minMana4 до maxMana
	p.Mana += restoredMana
	if p.Mana > p.ManaCap {
		p.Mana = p.ManaCap
	}

	// fmt.Printf("%s восстановил %d маны!\n",
	// 	p.Username, restoredMana)
}

// Heal лечит игрока за счет маны
func (p *Player) Heal(minHealth, maxHealth, minManaCost, maxManaCost int) bool {
	if p.Mana < 1 {
		// fmt.Printf("%s не хватает маны для лечения!\n", p.Username)
		return false
	}

	healAmount := rand.Intn(maxHealth-minHealth+1) + minHealth     // Восстановление от minHealth до 7
	manaCost := rand.Intn(maxManaCost-minManaCost+1) + minManaCost // Стоимость от minManaCost до 2 маны
	if p.Mana < maxManaCost {
		healAmount = healAmount / 2
		p.Health += healAmount
		if p.Health > p.HealthCap {
			p.Health = p.HealthCap
		}
		p.Mana = 0
	} else {
		p.Health += healAmount
		if p.Health > p.HealthCap {
			p.Health = p.HealthCap
		}
		p.Mana -= manaCost
	}

	// fmt.Printf("%s использовал лечение: +%d HP (-%d маны)\n",
	// 	p.Username, healAmount, manaCost)

	if p.Health > 50 {
		p.Health = 50
	}

	return true
}

// Проверяет жив ли игрок
func (p *Player) IsAlive() bool {
	return p.Health > 0
}
