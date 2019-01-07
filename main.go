package main

import (
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Client struct {
	BaseURL   *url.URL
	UserAgent string

	httpClient *http.Client
}

var Goal int
var debug bool

func main() {
	var showBelow = flag.Bool("below", false, "show characters below goal")
	var showAbove = flag.Bool("above", false, "show characters above goal")
	flag.BoolVar(&debug, "debug", false, "debug")
	flag.IntVar(&Goal, "goal", 38, "Azerite neck level goal")
	flag.Parse()
	cfgFile := "config.json"

	// read the configuration
	file, err := os.Open(cfgFile)
	if err != nil {
		fmt.Println("Could not open configuration file", cfgFile)
	}

	jsondec := json.NewDecoder(file)
	config := Configuration{}
	err = jsondec.Decode(&config)
	if err != nil {
		fmt.Println("Could not decode configuration file", cfgFile)
	}

	//Get guild members with specified rank
	GetGuildMembers(config.Realm, config.Guild, &config.GuildMembers, debug)
	for _, member := range config.GuildMembers.Members {
		for _, v := range config.GuildRanks {
			if v == member.Rank {
				config.Characters = append(config.Characters, &Character{
					Name:  member.Character.Name,
					Realm: member.Character.Realm,
				})
			}
		}

	}
	// make the channels, get the time, launch the goroutines
	characterCount := len(config.Characters)
	fmt.Println("Main :", characterCount, "accounts to collect data from in config")
	characters := make(chan *Character, characterCount)
	done := make(chan bool, characterCount)

	fmt.Println("Main : Submitting job to workers")

	for i, _ := range config.Characters {
		go worker(i, config, characters, done)
	}

	for _, character := range config.Characters {
		characters <- character

	}
	close(characters)

	for i := 0; i < characterCount; i++ {
		<-done
	}

	y := strings.Repeat("#", 100)
	fmt.Println(fmt.Sprintf("\n%s\n", y))
	fmt.Println("\t Goal", Goal, "\n")
	for _, character := range config.Characters {
		if Goal <= character.Items.Neck.AzeriteItem.AzeriteLevel && *showAbove {
			fmt.Println("Achieved goal:", character.Name, character.Items.Neck.AzeriteItem.AzeriteLevel)
		}
		if Goal >= character.Items.Neck.AzeriteItem.AzeriteLevel && *showBelow {
			fmt.Println("Below goal:", character.Name, character.Items.Neck.AzeriteItem.AzeriteLevel)
		}

	}

}

func (character *Character) QueryCharacter() {
	GetCharacter("GET", character.Realm, character.Name, character, debug)
}

func worker(id int, config Configuration, characters <-chan *Character, done chan<- bool) {
	for character := range characters {
		fmt.Println("Worker", id, ": Received job", character.Name)
		character.QueryCharacter()
		done <- true
	}
}

func GetGuildMembers(realm string, guild string, target interface{}, debug bool) error {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	url := "https://eu.api.blizzard.com/wow/guild/" + realm + "/" + guild + "?fields=members&locale=en_US&access_token=USHjtIcb3fA47RBFc7f431Q1fCol2HzSUU"
	if debug == true {
		fmt.Println("URL:>", url)
	}

	req, err := http.NewRequest("GET", url, nil)

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(&target)

}

func GetCharacter(method string, realm string, char string, target interface{}, debug bool) error {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	url := "https://eu.api.blizzard.com/wow/character/" + realm + "/" + char + "?fields=items&locale=en_US&access_token=USHjtIcb3fA47RBFc7f431Q1fCol2HzSUU"
	if debug == true {
		fmt.Println("URL:>", url)
	}

	req, err := http.NewRequest("GET", url, nil)

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(&target)

}
