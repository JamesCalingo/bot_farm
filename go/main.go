package main

import (
	"gofer_gopher/bot"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func getToken(key string) string {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal(err)
	}
	return os.Getenv(key)
}

func main() {
	// _, err := bot.AddChallenge("https://codingchallenges.fyi/challenges/challenge-password")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	bot.Token = getToken("GO")
	bot.Run()
}
