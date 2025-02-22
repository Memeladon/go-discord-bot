package constants

const CommandPrefix string = "))"

const CinemaRole string = "Cinema"

var ValidationRules = map[string]func([]string) bool{
	"": func(args []string) bool {
		return true
	},
	"game": func(args []string) bool {
		// game [username]
		return len(args) == 1
	},
	"cinema": func(args []string) bool {
		// cinema [title] [description] [cinemaRoleMention]
		return len(args) == 3
	},
	"question": func(args []string) bool {
		return true
	},
}

var DontUnderstandAnswers = [...]string{
	"не понял",
	"?",
	"а?",
}

var HostWhitelist = [...]string{
	"www.kinopoisk.ru",
}
