package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type SteamAPIResult struct {
	SteamID int `json:"steamid"`
	Success int `json:"success"`
}

func main() {

	steamAPIKey := "APIKeyHere"
	// var openDotaAPIKey string = ""

	fmt.Println("   ___  ____  _________     ___  ___  ____  __________   ____\n" +
		"  / _ \\/ __ \\/_  __/ _ |   / _ \\/ _ \\/ __ \\/ __/  _/ /  / __/\n" +
		" / // / /_/ / / / / __ |  / ___/ , _/ /_/ / _/_/ // /__/ _/  \n" +
		"/____/\\____/ /_/ /_/ |_| /_/  /_/|_|\\____/_/ /___/____/___/  \n" +
		"                                                             ")

	myUsername := getUserName()

	fmt.Println("I love Dota and my username is " + myUsername + "!")
	fmt.Println(getSteamID(steamAPIKey, myUsername))

}

func getUserName() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter your Steam Username: ")
	username, _ := reader.ReadString('\n')
	trimmedUsername := strings.TrimSuffix(username, "\n")
	return trimmedUsername
}

func getSteamID(steamAPIKey string, username string) int {
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

	var steamAPIResult SteamAPIResult

	json.Unmarshal(responseData, &steamAPIResult)

	fmt.Println("ID: " + string(steamAPIResult.SteamID))

	response := steamAPIResult.SteamID

	return response
}
