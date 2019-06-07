package opendota

import (
	"fmt"

	"github.com/richardziegler/go-dota/rest"
)

// WinsLosses ... explains itself
type WinsLosses struct {
	Win  int `json:"win"`
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

	var winsLosses WinsLosses
	rest.Get(openDotaUrl, winsLosses)

	tGames := winsLosses.Win + winsLosses.Lose
	aGames := int((float32(winsLosses.Win) / float32(tGames)) * 100)

	return winsLosses.Win, winsLosses.Lose, aGames
}

func GetPlayerProfileName(steamID string, openDotaKey string) string {
	openDotaUrl := fmt.Sprintf("https://api.opendota.com/api/players/%s?api_key=%s", steamID, openDotaKey)

	var playerProfile OdProfile
	rest.Get(openDotaUrl, playerProfile)

	return playerProfile.OdPData.PersonaName
}
