package storage

//#region Header
import (
	. "../models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	Token    string   `json:"token"`
	GitHub   string   `json:"github,omitempty"`
	Firebase Firebase `json:"firebase,omitempty"`
	User     bool     `json:"user"`
	Prefix   string   `json:"prefix"`
}

var Data Config

const (
	CONFIG_NAME string = "config.json"
)

//#endregion

func Load() {
	data, err := ioutil.ReadFile(CONFIG_NAME)
	if err != nil {
		log.Panic(err)
	}
	err = json.Unmarshal(data, &Data)
	if err != nil {
		log.Panic(err)
	}
	firebase, err := json.Marshal(&Data.Firebase)
	if err != nil {
		fmt.Println(err)
	} else {
		err = loadFirestore(firebase)
		if err != nil {
			fmt.Println(err)
		}
	}
	err = os.Setenv("TOKEN", Data.Token)
	if err != nil {
		fmt.Println(err)
	}
	err = os.Setenv("PREFIX", Data.Prefix)
	if err != nil {
		fmt.Println(err)
	}
	if len(Data.GitHub) == 0 {
		fmt.Println("github token is empty, github commands disabled")
	} else {
		err = os.Setenv("GITHUB", Data.GitHub)
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println("Config loaded")
}
