package cards

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
)

func getHomeDir() string {
	usr, err := user.Current()
	if err != nil {
		fmt.Println(err)
	}
	return usr.HomeDir + "/.greetctl"
}

// cardCmd represents the card command
func CreateAndSaveCard(config Config) {
	cardID := config.CardID
	occasion := config.Occasion
	lang := config.Language
	name := config.UserName

	homeDir := getHomeDir()
	if _, err := os.Stat(homeDir); os.IsNotExist(err) {
		os.Mkdir(homeDir, 0700)
	}

	file, _ := os.Create(homeDir + "/" + cardID)

	defer file.Close()

	var greeting string

	if occasion == newYear && lang == langEN {
		greeting = newYearGreetingENG + name
	} else if occasion == newYear && lang == langFR {
		greeting = newYearGreetingFR + name
	} else if occasion == txGiving && lang == langEN {
		greeting = txDayGreetingENG + name
	} else if occasion == txGiving && lang == langFR {
		greeting = txDayGreetingFR + name
	} else if occasion == bDay && lang == langEN {
		greeting = bDayGreetingENG + name
	} else {
		greeting = bDayGreetingFR + name
	}
	file.WriteString(greeting)
	fmt.Println("card with id [" + cardID + "] has been created")
}

// Fetch card's data
func FetchCardByID(id string) {
	data, _ := ioutil.ReadFile(getHomeDir() + "/" + id)
	fmt.Printf("Data: %s\n", data)
}

func FetchAllCards() {
	homeDir := getHomeDir()
	files, err := ioutil.ReadDir(homeDir)
	if err != nil {
		fmt.Println(err)
	}

	for _, f := range files {
		data, _ := ioutil.ReadFile(homeDir + "/" + f.Name())
		fmt.Printf("%s\n", data)
	}
}
