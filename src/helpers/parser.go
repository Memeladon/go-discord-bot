package helpers

import (
	"errors"
	"go-bot/src/constants"
	"strings"
)

// Cтруктура команды
type Command struct {
	Prefix   string
	Keyword  string
	Args     []string
	RawInput string
}

// Парсер команд
type Parser struct {
	prefix string
}

func NewParser(prefix string) *Parser {
	return &Parser{prefix: prefix}
}

// Parse парсит строку и возвращает структуру команды
func (p *Parser) Parse(input string) (*Command, error) {
	if !strings.HasPrefix(input, p.prefix) {
		return nil, errors.New("неверный префикс команды")
	}

	cmdStr := strings.TrimSpace(input[len(p.prefix):])
	parts := strings.Split(cmdStr, " ")

	if len(parts) < 1 {
		return nil, errors.New("отсутствует ключевое слово команды")
	}

	keyword := parts[0]
	args := parts[:]

	return &Command{
		Prefix:   p.prefix,
		Keyword:  keyword,
		Args:     args,
		RawInput: input,
	}, nil
}

// Validate проверяет валидность команды согласно правилам
func (c *Command) Validate(rules map[string]func([]string) bool) error {
	validator, exists := rules[c.Keyword]
	if !exists {
		return errors.New("неизвестная команда")
	}

	if !validator(c.Args) {
		return errors.New("неверные аргументы для команды")
	}

	return nil
}

// DefaultParser создает парсер с настройками по умолчанию
func DefaultParser() *Parser {
	return NewParser(constants.CommandPrefix)
}

// ValidateCommand проверяет команду с использованием стандартных правил
func ValidateCommand(cmd *Command) error {
	return cmd.Validate(constants.ValidationRules)
}
