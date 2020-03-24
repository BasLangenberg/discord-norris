package main

import (
	"github.com/BasLangenberg/discord-norris/internal/icndb"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"syscall"
)

const (
	version = "v0.1.0"
)

func main(){
	log.Printf("starting discord-norris %v\n", version)
	signalchan := make(chan os.Signal, 1)
	signal.Notify(signalchan, syscall.SIGINT, syscall.SIGTERM)

	dg, err := discordgo.New("Bot " + os.Getenv("DISCORD_BOT_KEY"))
	if err != nil {
		log.Printf("Unable to connect to discord: %v\n", err)
		os.Exit(1)
	}

	dg.AddHandler(responseWithQuote)

	err = dg.Open()
	if err != nil {
		log.Printf("unable to start bot: %v", err)
	}

	log.Println("Bot initialized and running, press CTRL+C to stop")

	for {
		select {
		case <- signalchan:
			log.Println("Terminating...")
			dg.Close()
			os.Exit(0)
		}
	}
}

func responseWithQuote(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!chuck" {
		quote, err := icndb.GetRandomQuote()
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Can't get a quote, please message @commissarbas who is supposed to maintain this bot")
		}
		s.ChannelMessageSend(m.ChannelID, quote)
	}
}