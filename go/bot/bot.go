package bot

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var Token string

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Run() {
	discord, err := discordgo.New("Bot " + Token)
	checkError(err)

	discord.AddHandler(newMessage)

	discord.Open()
	defer discord.Close()

	fmt.Println("Gofer Gopher Online")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

func newMessage(discord *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == discord.State.User.ID {
		return
	}

	switch {
	case strings.Contains(message.Content, "&hello"):
		greeting := "Hello " + message.Author.Username
		discord.ChannelMessageSend(message.ChannelID, greeting)
	case strings.Contains(message.Content, "&quote"):
		discord.ChannelMessageSend(message.ChannelID, GetQuote())
	case strings.Contains(message.Content, "&challenge"):
		discord.ChannelMessageSend(message.ChannelID, GetRandomChallenge("../challenges.json"))
	case strings.Contains(message.Content, "&list"):
		discord.ChannelMessageSend(message.ChannelID, GetChallengeList("../challenges.json"))
	case strings.Contains(message.Content, "&add"):
		_, url, _ := strings.Cut(message.Content, " ")
		newChallenge, err := AddChallenge(url)
		if err != nil {
			discord.ChannelMessageSend(message.ChannelID, "This is not a valid URL.")
			break
		}
		added := "$add" + " " + newChallenge.URL
		discord.ChannelMessageSend(message.ChannelID, added)
	}
}
