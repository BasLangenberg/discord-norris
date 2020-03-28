package giphy

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
)

func GetRandomChuckGifDownSizedLarge() (string, error) {
	var gif GiphyResponse

	if err := callGiphy(&gif); err != nil {
		return "", fmt.Errorf("unable to retrieve gif from giphy: %v", err)
	}


	i := rand.Intn(len(gif.Gif))
	return gif.Gif[i].Images.DownsizedLarge.URL, nil
}

func callGiphy(response *GiphyResponse) error {
	endpoint := "https://api.giphy.com/v1/gifs/search"
	apikey := os.Getenv("GIPHY_API")
	search := "chuck+norris"

	if apikey == "" {
		return fmt.Errorf("api key for giphy not set")
	}

	url := fmt.Sprint(endpoint + "?api_key=" + apikey + "&q=" + search + "&limit=25&offset=0&"+
		"rating=R&lang=en")
	req, err := http.Get(url)

	if err != nil {
		return fmt.Errorf("unable to call giphy endpoint")
	}

	if req.StatusCode < 200 || req.StatusCode > 299 {
		return fmt.Errorf("Error: HTTP %v returned", req.StatusCode)
	}

	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		return fmt.Errorf("unable to read body")
	}

	err = json.Unmarshal(body, &response)

	if err != nil {
		fmt.Errorf("unable to parse json")
	}

	return nil
}