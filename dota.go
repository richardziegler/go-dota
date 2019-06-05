package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/richardziegler/go-dota/opendota"
	"github.com/richardziegler/go-dota/steam"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

var steamAPIKey = "xxx"
var openDotaAPIKey = "xxx"

func main() {

	handleRequests()

	myUsername := getUserName()

	sID := strconv.FormatInt(steam.GetSteamID(steamAPIKey, myUsername), 10)
	displayName := opendota.GetPlayerProfileName(sID, openDotaAPIKey)
	fmt.Printf("Your username: %s\n", displayName)

}

func getUserName() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter your Steam Username: ")
	username, _ := reader.ReadString('\n')
	trimmedUsername := strings.TrimSuffix(username, "\n")
	return trimmedUsername
}

func getWinsLosses(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["id"]
	sID := strconv.FormatInt(steam.GetSteamID(steamAPIKey, username), 10)
	win, loss, average := opendota.GetWinsAndLosses(sID, openDotaAPIKey)
	json.NewEncoder(w).Encode(opendota.WinsLosses{win, loss, average})
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/getWinsLosses/{id}", getWinsLosses)
	log.Fatal(http.ListenAndServe(":1010", myRouter))
}
