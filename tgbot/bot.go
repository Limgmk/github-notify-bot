package tgbot

import (
	"github-notify-bot/model"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var bot *tgbotapi.BotAPI

func init() {
	token := model.GetConfig().TelegramBotToken
	bot, _ = tgbotapi.NewBotAPI(token)
}

func GetBot() *tgbotapi.BotAPI {
	return bot
}
