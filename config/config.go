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
}

var Global = &Config{}
var Mail = &(Global.Email)

func LoadConfig(filePath string) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	if err := toml.Unmarshal(content, Global); err != nil {
		panic(err)
	}
}
