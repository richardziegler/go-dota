package steam

import (
	"fmt"
	"log"
	"strconv"

	"github.com/richardziegler/go-dota/rest"
)

// R ... top level response structure
type R struct {
	APIResult APIResult `json:"response"`
}

// APIResult ... lower level data
type APIResult struct {
	SteamID string `json:"steamid"`
	Success int    `json:"success"`
}

func GetSteamID(steamAPIKey string, username string) int64 {
	steamAPIUrl := fmt.Sprintf("http://api.steampowered.com/ISteamUser/ResolveVanityURL/v0001/?key=%s&vanityurl=%s", steamAPIKey, username)

	var steamAPIData R
	rest.Get(steamAPIUrl, steamAPIData)

	uID64 := steamAPIData.APIResult.SteamID
	steamID64, err := strconv.ParseInt(uID64[3:len(uID64)], 0, 64)
	if err != nil {
		log.Fatal(err)
	}
	// Magic Number to get Steam32 ID
	sID := steamID64 - 61197960265728

	return sID
}
