package icndb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ICNDBQuoteResponse struct {
	Type  string `json:"type"`
	Value struct {
		ID         int           `json:"id"`
		Joke       string        `json:"joke"`
		Categories []interface{} `json:"categories"`
	} `json:"value"`
}

func GetRandomQuote() (string, error){
	resp, err := http.Get("https://api.icndb.com/jokes/random")
	if err != nil {
		return "", fmt.Errorf("failed to get a response: %v", err)
	}

	defer resp.Body.Close()

	payload, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to parse response: %v", err)
	}

	var quote ICNDBQuoteResponse
	err = json.Unmarshal(payload, &quote)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal payload: %v", err)
	}

	return quote.Value.Joke, nil
}
