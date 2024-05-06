package bot

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"slices"
	"strings"

	"github.com/bwmarrin/discordgo"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
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

var game Game
var validRoles = []string{"medic", "detective", "tanner", "mafioso"}

func contains(list []string, match string) bool {
	for _, elem := range list {
		if strings.EqualFold(elem, match) {
			return true
		}
	}
	return false
}

func countMafia(s []string) int {
	var count int
	for _, role := range s {
		if role == "mafioso" {
			count++
		}
	}
	return count
}

func printRoles(s []string) string {
	output := "CURRENT ROLES:\n"
	for _, role := range s {
		caser := cases.Title(language.AmericanEnglish)
		titled := caser.String(role)
		output += fmt.Sprintf("%s\n", titled)
	}
	return output
}

func newMessage(discord *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == discord.State.User.ID {
		return
	}

	switch {
	case strings.EqualFold(message.Content, "!openlobby") && game.isRunning:
		discord.ChannelMessageSend(message.ChannelID, "The lobby is already open")

	case strings.EqualFold(message.Content, "!openlobby"):
		game.toggleLobby()
		if !contains(game.roles, "mafioso") {
			game.roles = append(game.roles, "mafioso")
		}
		discord.ChannelMessageSend(message.ChannelID, "Starting a new game of Mafia. Type \"!join\" to join in.")

	case (strings.EqualFold(message.Content, "!join") || strings.EqualFold(message.Content, "!start") || strings.HasPrefix(strings.ToLower(message.Content), "!add")) && !game.isRunning:
		discord.ChannelMessageSend(message.ChannelID, "There is no active game going on right now. use !openlobby to start a game.")

	case strings.EqualFold(message.Content, "!join"):
		if contains(game.players, message.Author.ID) {
			discord.ChannelMessageSend(message.ChannelID, "You're already in the game!")
			return
		}
		discord.ChannelMessageSend(message.ChannelID, message.Author.Username+" has joined.")
		game.joinGame(message.Author.ID)

	case strings.HasPrefix(strings.ToLower(message.Content), "!add"):
		newRole := strings.TrimPrefix(strings.ToLower(message.Content), "!add ")
		if !contains(validRoles, newRole) {
			discord.ChannelMessageSend(message.ChannelID, "This role cannot be added.")
			return
		}
		game.roles = append(game.roles, newRole)
		discord.ChannelMessageSend(message.ChannelID, "Role added!")
		discord.ChannelMessageSend(message.ChannelID, printRoles(game.roles))

	case strings.HasPrefix(strings.ToLower(message.Content), "!remove"):
		role := strings.TrimPrefix(strings.ToLower(message.Content), "!remove ")
		if !contains(game.roles, role) {
			discord.ChannelMessageSend(message.ChannelID, "This role is not in this game.")
			return
		}

		if countMafia(game.roles) == 1 && role == "mafioso" {
			discord.ChannelMessageSend(message.ChannelID, "There needs to be at least one mafioso in the game.")
			return
		}
		i := slices.Index(game.roles, role)
		game.roles = slices.Delete(game.roles, i, i+1)
		discord.ChannelMessageSend(message.ChannelID, "Role deleted.")
		discord.ChannelMessageSend(message.ChannelID, printRoles(game.roles))

	case strings.EqualFold(message.Content, "!start"):
		if len(game.players) == 0 {
			return
		}
		if len(game.roles) > len(game.players) {
			discord.ChannelMessageSend(message.ChannelID, "You have too many roles for not enough players.")
			return
		}
		if len(game.players) == countMafia(game.roles) {
			discord.ChannelMessageSend(message.ChannelID, "EVERYONE'S A MAFIOSO. Reduce the mafia")
			return
		}
		for len(game.roles) < len(game.players) {
			game.roles = append(game.roles, "villager")
		}
		for _, id := range game.players {
			i := rand.Intn(len(game.roles))
			role := game.roles[i]
			output := fmt.Sprintf("You are a %s.\n", role)
			switch role {
			case "mafioso":
				output += "\"Unalive\" the villagers at night until only mafia remain."
			case "villager":
				output += "Find the mafia before they take over."
			case "detective":
				output += "Investigate one player each night to see if they're good or bad."
			case "medic":
				output += "Protect one player from the mafia each night. You may attempt to protect yourself, but only once."
			case "tanner":
				output += "Convince the villagers to vote you out, and you win!\nNOTE: You do NOT win if you are \"unalived\" by the mafia."
			}
			dm, dmErr := discord.UserChannelCreate(id)
			checkError(dmErr)
			discord.ChannelMessageSend(dm.ID, output)
			game.roles = slices.Delete(game.roles, i, i+1)
		}
		game.start()
		fmt.Println(game)
	}

}
