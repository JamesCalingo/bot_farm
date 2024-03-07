package bot

import (
	"encoding/json"
	"io"
	"net/http"
)

type Quote struct {
	Q string `json:"q"`
	A string `json:"a"`
}

func GetQuote() string {
	var quote []Quote
	res, _ := http.Get("https://zenquotes.io/api/random")
	data, _ := io.ReadAll(res.Body)
	json.Unmarshal(data, &quote)

	return "\"" + quote[0].Q + "\"\n-" + quote[0].A
}
