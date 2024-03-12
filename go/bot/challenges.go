package bot

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Challenges []Challenge

type Challenge struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func Open(fileName string) Challenges {
	var challenges Challenges
	file, fileErr := os.Open(fileName)
	if fileErr != nil {
		log.Fatal(fileErr)
	}
	data, _ := io.ReadAll(file)
	json.Unmarshal(data, &challenges)

	return challenges
}

func GetRandomChallenge(list Challenges) string {
	index := rand.Intn(len(list))
	randomChallenge := list[index]
	return randomChallenge.Name + ": " + randomChallenge.URL
}

func GetChallengeList(list Challenges) string {
	var output string
	for i := 0; i < len(list); i++ {
		output += list[i].Name + ": " + list[i].URL + "\n"
	}
	return output
}

func AddChallenge(url string) (Challenge, error) {
	var newChallenge Challenge
	if !strings.Contains(url, "https://codingchallenges.fyi/challenges/") {
		return newChallenge, errors.New("invalid url")
	}
	res, err := http.Get(url)
	if err != nil {
		return newChallenge, errors.New("error on get request")
	}
	if res.StatusCode != 200 {
		return newChallenge, errors.New("status code not 200")
	}
	defer res.Body.Close()
	website, _ := goquery.NewDocumentFromReader(res.Body)
	title := website.Find("Title").Text()
	name, _, _ := strings.Cut(title, " |")
	newChallenge.Name = name
	newChallenge.URL = url
	return newChallenge, nil
}
