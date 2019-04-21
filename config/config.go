package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	Token    string   `json:"token"`
	Prefix   string   `json:"prefix"`
	Commands []string `json:"commands"`
}

var Data Config

const (
	CONFIG_NAME    string      = "config.json"
	OS_PERMISSIONS os.FileMode = 0644
)

func LoadConfig() {
	data, err := ioutil.ReadFile(CONFIG_NAME)
	if err != nil {
		log.Panic(err)
	}
	err = json.Unmarshal(data, &Data)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("Loaded", len(Data.Commands), "commands")
}

func SaveConfig() {
	data, err := json.Marshal(&Data)
	if err != nil {
		log.Panic(err)
	}
	err = ioutil.WriteFile(CONFIG_NAME, data, OS_PERMISSIONS)
	if err != nil {
		log.Panic(err)
	}
}
