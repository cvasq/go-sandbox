package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type sshServer struct {
	Host     string `json:"host"`
	Username string `json:"username"`
	Key      string `json:"key"`
	URLMatch string `json:"urlMatch"`
}

func LoadConfiguration(file string) []sshServer {
	var config []sshServer
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}

func main() {

	config := LoadConfiguration("./config.json")

	for _, v := range config {
		fmt.Println(v.Host)
		fmt.Println(v.URLMatch)

	}
}
