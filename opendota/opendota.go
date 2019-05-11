package opendota

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// WinsLosses ... explains itself
type WinsLosses struct {
	Win int `json:"win"`
	Lose int `json:"lose"`
}

type OdProfile struct {
	OdPData OdPData `json:"profile""`
}

type OdPData struct {
	PersonaName string `json:"personaname""`
}

func GetWinsAndLosses(steamID string, openDotaKey string) (int, int, int) {
	openDotaUrl := fmt.Sprintf("https://api.opendota.com/api/players/%s/wl?api_key=%s", steamID, openDotaKey)

	openDotaResponse, err := http.Get(openDotaUrl)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(openDotaResponse.Body)
	if err != nil {
		log.Fatal(err)
	}

	var winsLosses WinsLosses
	json.Unmarshal(responseData, &winsLosses)

	tGames := winsLosses.Win + winsLosses.Lose
	aGames := int((float32(winsLosses.Win) / float32(tGames)) * 100)

	return winsLosses.Win, winsLosses.Lose, aGames
}

func GetPlayerProfileName(steamID string, openDotaKey string) string {
	openDotaUrl := fmt.Sprintf("https://api.opendota.com/api/players/%s?api_key=%s", steamID, openDotaKey)

	openDotaResponse, err := http.Get(openDotaUrl)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(openDotaResponse.Body)
	if err != nil {
		log.Fatal(err)
	}

	var playerProfile OdProfile
	json.Unmarshal(responseData, &playerProfile)

	return playerProfile.OdPData.PersonaName
}