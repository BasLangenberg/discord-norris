package main

import (
	"github.com/BasLangenberg/discord-norris/internal/giphy"
	"github.com/BasLangenberg/discord-norris/internal/icndb"
	mcsrvstat_us "github.com/BasLangenberg/discord-norris/internal/mcsrvstat-us"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
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

	go checkForMySon(dg)

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

	if strings.Contains(strings.ToLower(m.Content), "!chuck") {
		quote, qerr := icndb.GetRandomQuote()
		gif, gerr := giphy.GetRandomChuckGifDownSizedLarge()
		if qerr != nil {
			s.ChannelMessageSend(m.ChannelID, "Can't get a quote, please message @commissarbas who is supposed to maintain this bot")
		}
		if gerr != nil {
			s.ChannelMessageSend(m.ChannelID, "Can't get a gif, please message @commissarbas who is supposed to maintain this bot")
		}
		embed := &discordgo.MessageEmbed{
			Author:      &discordgo.MessageEmbedAuthor{},
			Color:       0x00ff00, // Green
			Description: quote,
			Image: &discordgo.MessageEmbedImage{
				URL: gif,
			},
			Timestamp: time.Now().Format(time.RFC3339), // Discord wants ISO8601; RFC3339 is an extension of ISO8601 and should be completely compatible.
			Title:     "Chuck Norris Quote",
		}

		_, err := s.ChannelMessageSendEmbed(m.ChannelID, embed)

		if err != nil {
			log.Printf("error publishing message: %v", err)
		}

	}
}

func checkForMySon(s *discordgo.Session) {
	var cache []string
	for {
		time.Sleep(1 * time.Minute)
		log.Println("Checking for specific username")

		online, err := mcsrvstat_us.GetOnlinePlayers()

		if err != nil {
			log.Printf("Unable to get online playes: %v\n", err)
		}

		// Check
		for _, gamer := range online {
			if strings.Contains(gamer, "Sebe") {
				// Discord meuk
				if !isStringInStringSlice(cache, gamer) {
					s.ChannelMessageSend("692063079483047989", "Sébe just came online")
				}
			}
			if strings.Contains(gamer, "Bas") {
				if !isStringInStringSlice(cache, gamer) {
					s.ChannelMessageSend("692063079483047989", "Bas just came online")
				}
			}
		}

		cache = online
		log.Printf("online users: %v", cache)

	}
}

func isStringInStringSlice(cache []string, gamer string) bool {
	for _, user := range cache {
		if user == gamer {
			return true
		}
	}

	return false
}