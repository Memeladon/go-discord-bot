package helpers

import (
	"fmt"
	"go-bot/src/constants"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var (
	// небольшой кэш, чтобы не искать роль каждый раз
	rolesMap = make(map[string]*discordgo.Role) //RoleName + GuildId -> Role
)

func ParseCommand(m *discordgo.MessageCreate) []string {
	withoutPrefix := m.Content[len(constants.CommandPrefix):]
	return strings.Split(withoutPrefix, " ")
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
