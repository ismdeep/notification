package config

import (
	"github.com/BurntSushi/toml"
	"io/ioutil"
)

// Config 配置结构体
type Config struct {
	Bind   string
	Secret string
	Email  struct {
		Host     string
		Port     int
		Username string
		Password string
	} `json:"email"`
	JWT struct {
		Expire string
		Key    string
	} `json:"jwt"`
	MySQL struct {
		DSN string
	} `json:"mysql"`
}

var Global = &Config{}
var Mail = &(Global.Email)
var MySQL = &(Global.MySQL)
var JWT = &(Global.JWT)

func LoadConfig(filePath string) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	if err := toml.Unmarshal(content, Global); err != nil {
		panic(err)
	}
}
