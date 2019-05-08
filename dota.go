package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type R struct {
	SteamAPIResult SteamAPIResult `json:"response"`
}

type SteamAPIResult struct {
	SteamID string `json:"steamid"`
	Success int    `json:"success"`
}

func main() {

	SteamAPIKey := "APIKeyHere"
	// var openDotaAPIKey string = ""

	fmt.Println("   ___  ____  _________     ___  ___  ____  __________   ____\n" +
		"  / _ \\/ __ \\/_  __/ _ |   / _ \\/ _ \\/ __ \\/ __/  _/ /  / __/\n" +
		" / // / /_/ / / / / __ |  / ___/ , _/ /_/ / _/_/ // /__/ _/  \n" +
		"/____/\\____/ /_/ /_/ |_| /_/  /_/|_|\\____/_/ /___/____/___/  \n" +
		"                                                             ")

	myUsername := getUserName()

	fmt.Println("I love Dota and my username is " + myUsername + "!")
	getSteamID(SteamAPIKey, myUsername)

}

func getUserName() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter your Steam Username: ")
	username, _ := reader.ReadString('\n')
	trimmedUsername := strings.TrimSuffix(username, "\n")
	return trimmedUsername
}

func getSteamID(SteamAPIKey string, username string) int64 {
	steamAPIUrl := fmt.Sprintf("http://api.steampowered.com/ISteamUser/ResolveVanityURL/v0001/?key=%s&vanityurl=%s", SteamAPIKey, username)

	steamResponse, err := http.Get(steamAPIUrl)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(steamResponse.Body)
	if err != nil {
		log.Fatal(err)
	}

	var steamAPIResult R

	json.Unmarshal(responseData, &steamAPIResult)

	fmt.Println("ID: " + string(steamAPIResult.SteamAPIResult.SteamID))

	uID64 := steamAPIResult.SteamAPIResult.SteamID
	steamID64, err := strconv.ParseInt(uID64[3:len(uID64)], 0, 64)
	if err != nil {
		log.Fatal(err)
	}

	sID := steamID64 - 61197960265728

	fmt.Println(sID)
	return sID
}
