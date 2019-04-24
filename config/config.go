package config

//#region Header
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	Token  string `json:"token"`
	GitHub string `json:"github"`
	User   bool   `json:"user"`
	Prefix string `json:"prefix"`
}

var Data Config

const (
	CONFIG_NAME string = "config.json"
)

//#endregion

func LoadConfig() {
	data, err := ioutil.ReadFile(CONFIG_NAME)
	if err != nil {
		log.Panic(err)
	}
	err = json.Unmarshal(data, &Data)
	if err != nil {
		log.Panic(err)
	}
	err = os.Setenv("TOKEN", Data.Token)
	if err != nil {
		fmt.Println(err)
	}
	err = os.Setenv("PREFIX", Data.Prefix)
	if err != nil {
		fmt.Println(err)
	}
	err = os.Setenv("GITHUB", Data.GitHub)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Loaded config")
}
