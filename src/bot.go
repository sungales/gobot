package bot

import (
	"fmt"
	"os"
	"os/signal"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var BotToken string

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}

func Run() { 
	discord, err := discordgo.New("Bot " + BotToken)
	checkNilErr(err)

	discord.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}
		switch { 
			case strings.Contains(m.Content, "!help"):
				discord.ChannelMessageSend(m.ChannelID, "Available commands:\n!help - Show this help message\n!ping - Respond with Pong!")
			case strings.Contains(m.Content, "!ping"):
				discord.ChannelMessageSend(m.ChannelID, "Pong!")
		}
	})

	discord.Open()
	defer discord.Close()

	fmt.Println("running")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

