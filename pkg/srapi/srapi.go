package srapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func GetGameByID(id string) (*Game, error) {
	resp, err := http.Get(fmt.Sprintf("https://www.speedrun.com/api/v1/games/%s", id))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var respData struct {
		Data *Game
	}
	if err := json.Unmarshal(data, &respData); err != nil {
		return nil, err
	}
	return respData.Data, nil
}

func GetGameByName(query string) ([]*Game, error) {
	resp, err := http.Get(fmt.Sprintf("https://www.speedrun.com/api/v1/games?name=%s", url.QueryEscape(query)))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var respData struct {
		Data []*Game
	}
	if err := json.Unmarshal(data, &respData); err != nil {
		return nil, err
	}
	return respData.Data, nil
}
