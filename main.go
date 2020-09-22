package main

import (
	"github-notify-bot/model"
	"github-notify-bot/router"
	"log"
	"os"
)

func main() {
	config := model.GetConfig()

	err := model.InitSQLite()
	if err != nil {
		log.Println(err)
		return
	}

	listenAddr := config.BindAddress
	if len(os.Args) > 1 {
		listenAddr = os.Args[1]
	}

	r := router.InitRouter()
	if config.CertPath != "" && config.KeyPath != "" {
		err = r.RunTLS(listenAddr, config.CertPath, config.KeyPath)
	} else {
		err = r.Run(listenAddr)
	}
	if err != nil {
		log.Fatal(err)
	}
}
