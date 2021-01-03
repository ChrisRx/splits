package srapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Game struct {
	ID    string
	Names struct {
		International string
		Japanese      string
		Twitch        string
	}
	Abbreviation string
	Weblink      string
	Released     int
	Ruleset      struct {
		ShowMilliseconds    bool           `json:"show-milliseconds"`
		RequireVerification bool           `json:"require-verification"`
		RequireVideo        bool           `json:"require-video"`
		RunTimes            []TimingMethod `json:"run-times"`
		DefaultTime         TimingMethod   `json:"default-time"`
		EmulatorsAllowed    bool           `json:"emulators-allowed"`
	}
	Romhack bool
	Created *time.Time
	Assets  map[string]*AssetLink
	Links   []Link
}

func (g *Game) Name() string {
	return g.Names.International
}

type Category struct {
	ID      string
	Name    string
	Weblink string
	Type    string
	Rules   string
	Players struct {
		Type  string
		Value int
	}
	Miscellaneous bool
	Links         []Link
}

func (g *Game) Categories() ([]*Category, error) {
	resp, err := http.Get(fmt.Sprintf("https://www.speedrun.com/api/v1/games/%s/categories", g.ID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var respData struct {
		Data []*Category
	}
	if err := json.Unmarshal(data, &respData); err != nil {
		return nil, err
	}
	return respData.Data, nil
}

type TimingMethod string

const (
	TimingIngameTime           TimingMethod = "ingame"
	TimingRealtime             TimingMethod = "realtime"
	TimingRealtimeWithoutLoads TimingMethod = "realtime_noloads"
)

type Link struct {
	Rel string
	URI string
}

type AssetLink struct {
	URI           string
	Width, Height int
}
