package bot

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Quote struct {
	Q string `json:"q"`
	A string `json:"a"`
}

func GetQuote() string {
	//Our API returns a quote in an array, so in order to get the actual object, we need to put it in a slice
	var quote []Quote
	res, _ := http.Get("https://zenquotes.io/api/random")
	data, _ := io.ReadAll(res.Body)
	json.Unmarshal(data, &quote)
	return fmt.Sprintf("\"%s\"\n-%s", quote[0].Q, quote[0].A)

}
