package agents

import (
	"encoding/json"
	"net/http"
	"strings"
)

var (
	// latestChromeAgent is initialized by init.
	latestChromeAgent string
	// chromeVersion is initialized by init.
	chromeVersion string
	// shortChromeVersion is initialized by init.
	shortChromeVersion string
)

// init updates the latest Chrome agent and Chrome version.
func init() {
	resp, err := http.Get("https://cdn.jsdelivr.net/gh/jnrbsn/user-agents@master/user-agents.json")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var data []string
	if err = json.NewDecoder(resp.Body).Decode(&data); err != nil {
		panic(err)
	}

	latestChromeAgent = data[0]
	chromeVersion = strings.Split(strings.Split(latestChromeAgent, "Chrome/")[1], " ")[0]
	shortChromeVersion = strings.Split(chromeVersion, ".")[0]
}
