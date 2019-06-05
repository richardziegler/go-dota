package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/richardziegler/go-dota/opendota"
	"github.com/richardziegler/go-dota/steam"
)

func main() {

	var steamAPIKey = "xxx"
	var openDotaAPIKey = "xxx"

	fmt.Println("   ___  ____  _________     ___  ___  ____  __________   ____\n" +
		"  / _ \\/ __ \\/_  __/ _ |   / _ \\/ _ \\/ __ \\/ __/  _/ /  / __/\n" +
		" / // / /_/ / / / / __ |  / ___/ , _/ /_/ / _/_/ // /__/ _/  \n" +
		"/____/\\____/ /_/ /_/ |_| /_/  /_/|_|\\____/_/ /___/____/___/  \n" +
		"                                                             ")

	myUsername := getUserName()

	sID := strconv.FormatInt(steam.GetSteamID(steamAPIKey, myUsername), 10)
	wins, losses, winR := opendota.GetWinsAndLosses(sID, openDotaAPIKey)
	displayName := opendota.GetPlayerProfileName(sID, openDotaAPIKey)

	fmt.Println("=====================================================================================")
	fmt.Printf("Your username: %s\n", displayName)
	fmt.Printf("You have won %v games and lost %v games\n", wins, losses)
	fmt.Printf("Your win rate is %v%%\n", winR)
	fmt.Println("=====================================================================================")

}

func getUserName() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter your Steam Username: ")
	username, _ := reader.ReadString('\n')
	trimmedUsername := strings.TrimSuffix(username, "\n")
	return trimmedUsername
}
