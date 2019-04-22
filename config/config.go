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
	Token    string   `json:"token"`
	User     bool     `json:"user"`
	Prefix   string   `json:"prefix"`
	Commands []string `json:"commands"`
}

var Data Config

const (
	CONFIG_NAME    string      = "config.json"
	OS_PERMISSIONS os.FileMode = 0644
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
	fmt.Println("Loaded config")
}

// todo keep or remove
func SaveConfig() {
	data, err := json.Marshal(&Data)
	if err != nil {
		log.Panic(err)
	}
	err = ioutil.WriteFile(CONFIG_NAME, data, OS_PERMISSIONS)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("Config successfully saved")
}
