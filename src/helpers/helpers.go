package helpers

import (
	"fmt"
	"go-bot/src/constants"
	"math/rand"
	"net/url"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	// небольшой кэш, чтобы не искать роль каждый раз
	rolesMap = make(map[string]*discordgo.Role) //RoleName + GuildId -> Role
)

func GetRandomElement(arr []string) string {
	rng := rand.New(rand.NewSource(time.Now().Unix()))
	randomIndex := rng.Intn(len(arr))
	return arr[randomIndex]
}

func IsWhitelistedHost(u *url.URL) bool {
	for _, host := range constants.HostWhitelist {
		if u.Hostname() == host {
			return true
		}
	}

	return false
}

func ParseCommand(m *discordgo.MessageCreate) []string {
	parser := DefaultParser()

	fmt.Printf("Входная комманда: %s\n", m.Content)

	// Парсим команду
	parsedCommand, err := parser.Parse(m.Content)
	if err != nil {
		fmt.Printf("Ошибка парсинга комманды: %v\n", err)
	}

	// Валидируем команду с использованием правил
	err = ValidateCommand(parsedCommand)
	if err != nil {
		fmt.Printf("Ошибка валидации: %v\n", err)
	} else {
		fmt.Println("Команда валидна")
		fmt.Printf("Ключевое слово: %s\n", parsedCommand.Keyword)
		fmt.Printf("Аргументы: %v\n", parsedCommand.Args)
	}
	fmt.Println()

	return parsedCommand.Args
}

func FindRoleInGuildByName(roleName string, guildId string, session *discordgo.Session) (*discordgo.Role, error) {
	hash := roleName + guildId

	guild, _ := session.Guild(guildId)

	if cachedRole, found := rolesMap[hash]; found {
		return cachedRole, nil
	}

	for _, role := range guild.Roles {
		if role.Name == roleName {
			rolesMap[hash] = role
			return role, nil
		}
	}

	return nil, fmt.Errorf("на сервере нет роли %s", roleName)
}
