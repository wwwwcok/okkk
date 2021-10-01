package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type AppConfig struct {
	AppName    string `json:"app_name"`
	Port       string `json:"port"`
	Staticpath string `json:"static_path"`
	Mode       string `json:"mode"`
}

var ServConfig AppConfig

func InitConfig() *AppConfig {
	file, _ := os.Open("C:/Users/pc/go/src/goweb_iris/day5_project/util/config.json")
	decoder := json.NewDecoder(file)
	config := AppConfig{}
	err := decoder.Decode(&config)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(config)
	fmt.Println("ssssssEeeeeeeeeeeeeeeeeeeeeeeEEEEEEEEEEEE")
	return &config
}
