package steam

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
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

var steamAPIKey = "APIKeyHere"

func getSteamID(SteamAPIKey string, username string) int64 {
	steamAPIUrl := fmt.Sprintf("http://api.steampowered.com/ISteamUser/ResolveVanityURL/v0001/?key=%s&vanityurl=%s", steamAPIKey, username)

	steamResponse, err := http.Get(steamAPIUrl)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(steamResponse.Body)
	if err != nil {
		log.Fatal(err)
	}

	var steamAPIData R

	json.Unmarshal(responseData, &steamAPIResult)

	fmt.Println("ID: " + string(steamAPIData.APIResult.SteamID))

	uID64 := steamAPIData.APIResult.SteamID
	steamID64, err := strconv.ParseInt(uID64[3:len(uID64)], 0, 64)
	if err != nil {
		log.Fatal(err)
	}

	sID := steamID64 - 61197960265728

	fmt.Println(sID)
	return sID
}
