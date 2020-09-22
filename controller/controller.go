package controller

import (
	"fmt"
	"github-notify-bot/model"
	"github-notify-bot/tgbot"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func ParseGithubMessage(c *gin.Context) {
	message := new(model.GithubMessage)
	c.ShouldBindJSON(message)

	if message.HookId != 0 {
		log.Println("A new Repo was added")
		repo := new(model.Repository)
		repo.FullName = message.Repository.FullName
		model.CreateRepo(repo)
		c.String(http.StatusOK, "ok")
		return
	}

	includeFiles := ""
	for _, filename := range message.HeadCommit.Added {
		includeFiles = includeFiles + " " + filename
	}
	for _, filename := range message.HeadCommit.Removed {
		includeFiles = includeFiles + " " + filename
	}
	for _, filename := range message.HeadCommit.Modified {
		includeFiles = includeFiles + " " + filename
	}
	notificationText := fmt.Sprintf(`
新推送 !

-----------------------

Repository:[ %v ](https://github.com/%v)

Author:[ %v ](https://github.com/%v)

Email: %v

Description: %v

Include: %v

Time: %v
`, message.Repository.FullName, message.Repository.FullName, message.HeadCommit.Author.Name, message.HeadCommit.Author.Username, message.HeadCommit.Author.Email, message.HeadCommit.Message, includeFiles, message.HeadCommit.Timestamp)
	bot := tgbot.GetBot()
	repo := model.Repository{
		FullName: message.Repository.FullName,
	}
	subscriberList, err := model.FindSubscribersByRepo(&repo)
	if err != nil {
		log.Println(err)
		c.String(http.StatusOK, "ok")
		return
	}
	for _, subscriber := range subscriberList {
		notification := tgbotapi.NewMessage(subscriber.ChatId, notificationText)
		notification.ParseMode = tgbotapi.ModeMarkdown
		notification.DisableWebPagePreview = true
		_, err := bot.Send(notification)
		if err != nil {
			log.Println(err)
		}
	}
	c.String(http.StatusOK, "ok")
}

func AuthTelegramUser(c *gin.Context)  {
	update := new(tgbotapi.Update)
	c.ShouldBindJSON(update)
	message := update.Message
	if message != nil && message.IsCommand() {
		command := message.Command()
		switch command {
		case "start":
			startCommand(update)
		case "auth":
			authCommand(update)
		case "add":
			addCommand(update)
		case "del":
			delCommand(update)
		case "sub":
			sub(update)
		case "unsub":
			unsub(update)
		case "list":
			listCommand(update)
		case "all":
			allCommand(update)
		default:
			otherCommand(update)
		}
	}
	c.String(http.StatusOK, "ok")
}
