package bot

import (
	"encoding/json"
	"math/rand"
	// "fmt"
	"io"
	"log"
	"os"
)

type Challenges [] Challenge

type Challenge struct {
	Name string `json:"name"`
	URL string `json:"url"`
}

func open(fileName string) Challenges {
	var challenges Challenges
	file, fileErr := os.Open(fileName)
	if fileErr != nil {
		log.Fatal(fileErr)
	}
	data, _ := io.ReadAll(file)
	json.Unmarshal(data, &challenges)

	return challenges
}

func GetRandomChallenge(filename string) string {
	list := open(filename)
	index := rand.Intn(len(list))
	randomChallenge := list[index]
return randomChallenge.Name + ": " + randomChallenge.URL
}

func GetChallengeList(filename string) string {
	var output string
	list := open(filename)
	for i := 0; i < len(list); i ++ {
		output += list[i].Name + ": " + list[i].URL + "\n"
	} 
return output
}