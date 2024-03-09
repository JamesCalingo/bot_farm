package main

import (
	"fmt"
	"gofer_gopher/bot"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func getToken(key string) string {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal(err)
	}
	return os.Getenv(key)
}

func testAdd(url string) {

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	data, _ := io.ReadAll(res.Body)
	website := string(data)
	title := strings.Contains(website, "<title")
	fmt.Println(title)

}

func main() {
	// testAdd("https://codingchallenges.fyi/challenges/challenge-password-cracker")
	bot.Token = getToken("GO")
	bot.Run()
}
