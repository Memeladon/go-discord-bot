package helpers

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var (
	// небольшой кэш, чтобы не искать каждый раз роль
	rolesMap = make(map[string]map[string]*discordgo.Role) //RoleName -> (GuildId  -> Role)
)

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
	guild, _ := session.Guild(guildId)
	if cachedRoleInGuild, found := rolesMap[roleName]; found {
		if cachedRole, found := cachedRoleInGuild[guildId]; found {
			return cachedRole, nil
		}
	}

	for _, role := range guild.Roles {
		if role.Name == roleName {
			rolesMap[roleName] = map[string]*discordgo.Role{
				guildId: role,
			}
			return role, nil
		}
	}

	return nil, fmt.Errorf("на сервере нет роли %s", roleName)
}
