package constants

const CommandPrefix string = "))"

const CinemaRole string = "Cinema"

var ValidationRules = map[string]func([]string) bool{
	"game": func(args []string) bool {
		// game [username]
		return len(args) == 1
	},
	"cinema": func(args []string) bool {
		// cinema [title] [description] [cinemaRoleMention]
		return len(args) == 3
	},
}
