package main

import (
	"fmt"
	"go-bot/src/middleware"
	"go-bot/src/modules/cinema"
	"go-bot/src/modules/magicball"
	"go-bot/src/modules/smile"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	handlers := [...]func(s *discordgo.Session, m *discordgo.MessageCreate){
		magicball.Handler,
		smile.Handler,
		cinema.Handler,
	}

	for _, handler := range handlers {
		modifiedHandler := middleware.CheckCommandMiddleware(middleware.IgnoreSelfMiddleware(handler))
		dg.AddHandler(modifiedHandler)
	}

	dg.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}
