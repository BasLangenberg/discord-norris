package mcsrvstat_us

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type mcsrvstatusResponse struct {
	IP    string `json:"ip"`
	Port  int    `json:"port"`
	Debug struct {
		Ping          bool `json:"ping"`
		Query         bool `json:"query"`
		Srv           bool `json:"srv"`
		Querymismatch bool `json:"querymismatch"`
		Ipinsrv       bool `json:"ipinsrv"`
		Cnameinsrv    bool `json:"cnameinsrv"`
		Animatedmotd  bool `json:"animatedmotd"`
		Proxypipe     bool `json:"proxypipe"`
		Cachetime     int  `json:"cachetime"`
		Apiversion    int  `json:"apiversion"`
	} `json:"debug"`
	Motd struct {
		Raw   []string `json:"raw"`
		Clean []string `json:"clean"`
		HTML  []string `json:"html"`
	} `json:"motd"`
	Players struct {
		Online int      `json:"online"`
		Max    int      `json:"max"`
		List   []string `json:"list"`
	} `json:"players"`
	Version  string `json:"version"`
	Online   bool   `json:"online"`
	Protocol int    `json:"protocol"`
	Hostname string `json:"hostname"`
	Icon     string `json:"icon"`
}

func GetOnlinePlayers() ([]string, error) {
	var response mcsrvstatusResponse

	url := "https://api.mcsrvstat.us/2/mc.homecooked.nl"
	resp, err := http.Get(url)

	if err != nil {
		return []string{}, fmt.Errorf("Unable to query mcsrvstat.us: %v", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return []string{}, fmt.Errorf("Error: HTTP %v returned", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return []string{}, fmt.Errorf("Unable to read response: %v", err)
	}

	err = json.Unmarshal(body, &response)

	if err != nil {
		return []string{}, fmt.Errorf("Unable to unmarshal JSON: %v", err)
	}

	return response.Players.List, nil

}