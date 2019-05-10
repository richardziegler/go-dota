package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	
	"github.com/richardziegler/go-dota/steam"
)

func main() {

	var steamAPIKey = "APIKeyHere"
	// var openDotaAPIKey string = ""

	fmt.Println("   ___  ____  _________     ___  ___  ____  __________   ____\n" +
		"  / _ \\/ __ \\/_  __/ _ |   / _ \\/ _ \\/ __ \\/ __/  _/ /  / __/\n" +
		" / // / /_/ / / / / __ |  / ___/ , _/ /_/ / _/_/ // /__/ _/  \n" +
		"/____/\\____/ /_/ /_/ |_| /_/  /_/|_|\\____/_/ /___/____/___/  \n" +
		"                                                             ")

	myUsername := getUserName()

	fmt.Println("I love Dota and my username is " + myUsername + "!")
	steam.GetSteamID(steamAPIKey, myUsername)

}

func getUserName() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter your Steam Username: ")
	username, _ := reader.ReadString('\n')
	trimmedUsername := strings.TrimSuffix(username, "\n")
	return trimmedUsername
}
