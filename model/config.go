package model

import (
	"github-notify-bot/util"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// Config :  配置文件结构
type Config struct {
	TelegramBotToken	string	`yaml:"bot-token"`
	BindAddress			string	`yaml:"bind-address"`
	Secret				string	`yaml:"secret"`
	CertPath			string	`yaml:"cert-path"`
	KeyPath				string	`yaml:"key-path"`
}

var config *Config

func init() {
	// 解析配置文件
	config = new(Config)
	configBytes, _ := ioutil.ReadFile(util.GetCurrentDirectory() + "/conf/config.yaml")

	if err := yaml.Unmarshal(configBytes, config); err != nil {
		log.Fatalln(err)
	}
}

func GetConfig() *Config {
	return config
}