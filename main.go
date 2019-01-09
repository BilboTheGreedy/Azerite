package main

import (
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/negah/percent"
	"github.com/olekukonko/tablewriter"
)

type Client struct {
	BaseURL   *url.URL
	UserAgent string

	httpClient *http.Client
}

var Goal int
var debug bool
var region string
var token string

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
	region = config.Region
	GetToken(config.ClientID, config.ClientSecret, &config.Session)
	token = config.Session.AccessToken
	//Get guild members with specified rank
	GetGuildMembers(region, config.Realm, config.Guild, &config.GuildMembers, debug)
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

	fmt.Println(fmt.Sprintf("\n%s\n", strings.Repeat("#", 100)))
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Goal", "Character", "Neck", "% lvl"})
	table.SetFooter([]string{"", "", "Goal", strconv.Itoa(Goal)}) // Add Footer
	table.SetBorder(false)
	for _, character := range config.Characters {
		if Goal <= character.Items.Neck.AzeriteItem.AzeriteLevel && *showAbove {
			perc := percent.PercentOf(character.Items.Neck.AzeriteItem.AzeriteExperience, character.Items.Neck.AzeriteItem.AzeriteExperienceRemaining)
			pstr := fmt.Sprintf("%.1f", perc)
			data := []string{"Yes", character.Name, strconv.Itoa(character.Items.Neck.AzeriteItem.AzeriteLevel), pstr}

			table.Append(data)

			//fmt.Println("Goal: Yes\t", character.Name+"\t", character.Items.Neck.AzeriteItem.AzeriteLevel)
		}
		if Goal > character.Items.Neck.AzeriteItem.AzeriteLevel && *showBelow {
			perc := percent.PercentOf(character.Items.Neck.AzeriteItem.AzeriteExperience, character.Items.Neck.AzeriteItem.AzeriteExperienceRemaining)
			pstr := fmt.Sprintf("%.1f", perc)
			data := []string{"No", character.Name, strconv.Itoa(character.Items.Neck.AzeriteItem.AzeriteLevel), pstr}
			table.Append(data)
			//fmt.Println("Goal No\t:", character.Name+"\t", character.Items.Neck.AzeriteItem.AzeriteLevel)
		}

	}
	table.Render()

}

func (character *Character) QueryCharacter() {
	GetCharacter(region, character.Realm, character.Name, character, debug)
}

func worker(id int, config Configuration, characters <-chan *Character, done chan<- bool) {
	for character := range characters {
		if debug == true {
			fmt.Println("Worker", id, ": Received job", character.Name)
		}

		character.QueryCharacter()
		done <- true
	}
}

func GetGuildMembers(region string, realm string, guild string, target interface{}, debug bool) error {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	url := "https://" + region + ".api.blizzard.com/wow/guild/" + realm + "/" + guild + "?fields=members&locale=en_US&access_token=" + token
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

func GetCharacter(region string, realm string, char string, target interface{}, debug bool) error {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	url := "https://" + region + ".api.blizzard.com/wow/character/" + realm + "/" + char + "?fields=items&locale=en_US&access_token=" + token
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

func basicAuth(username string, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

//GetToken  auth token
func GetToken(ClientID string, ClientSecret string, target interface{}) error {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	if debug == true {
		fmt.Println("URL:>", "https://"+region+".battle.net/oauth/token")
	}

	values := url.Values{}
	values.Add("grant_type", "client_credentials")

	req, err := http.NewRequest("POST", "https://eu.battle.net/oauth/token", strings.NewReader(values.Encode()))
	req.Header.Add("Authorization", "Basic "+basicAuth(ClientID, ClientSecret))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{Transport: tr}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 && debug == true {
		fmt.Println("response Status:", resp.Status)

	}

	return json.NewDecoder(resp.Body).Decode(&target)

}
