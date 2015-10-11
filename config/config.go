package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Sitename    string `json:"sitename"`
	Fqdn        string `json:"fqdn"`
	Host        string `json:"host"`
	SessionName string `json:"sessionname"`
	AppPort     string `json:"appport"`
}

func Conf() *Config {
	var Conf = new(Config)
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	err := decoder.Decode(Conf)
	if err != nil {
		fmt.Println("error:", err)
	}
	return Conf
}
