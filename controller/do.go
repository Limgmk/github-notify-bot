package controller

import (
	"fmt"
	"github-notify-bot/model"
	"github-notify-bot/tgbot"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

// 其他命令处理
func otherCommand(update *tgbotapi.Update) {
	replyText := "无效的命令"
	sendText(update.Message.Chat.ID, replyText)
}

// 处理 start 命令
func startCommand(update *tgbotapi.Update) {
	replyText := `/auth secret 验证权限
/sub repository 订阅通知
/unsub repository 取消通知
/list 查看所有订阅
`
	sendText(update.Message.Chat.ID, replyText)
}

// 处理 auth 命令
func authCommand(update *tgbotapi.Update) {
	message := update.Message
	chatId := message.Chat.ID
	replyText := ""

	secret := message.CommandArguments()
	if secret == model.GetConfig().Secret {
		subscriber := new(model.Subscriber)
		subscriber.UserName = message.Chat.UserName
		subscriber.ChatId = chatId
		model.CreateSubscriber(subscriber)
		replyText = "鉴权通过"
	} else {
		replyText = "鉴权失败"
	}

	sendText(chatId, replyText)
}

// 处理 add 命令
func addCommand(update *tgbotapi.Update) {
	if !verifyId(update) {
		return
	}
	message := update.Message
	chatId := message.Chat.ID
	replyText := ""

	if message.CommandArguments() == "" {
		replyText = "请用仓库全名作为参数"
	} else {
		repo := new(model.Repository)
		repo.FullName = message.CommandArguments()
		model.CreateRepo(repo)

		replyText = "添加仓库成功"
	}

	sendText(chatId, replyText)
}

// 处理 del 命令
func delCommand(update *tgbotapi.Update) {
	if !verifyId(update) {
		return
	}
	message := update.Message
	chatId := message.Chat.ID
	replyText := ""

	if repo, _ := model.FindRepoByFullName(message.CommandArguments()); repo != nil {
		model.DeleteRepo(repo)
		replyText = "删除成功"
	} else {
		replyText = "该仓库未收录"
	}

	sendText(chatId, replyText)
}

// 处理 sub 命令
func sub(update *tgbotapi.Update) {
	if !verifyId(update) {
		return
	}
	message := update.Message
	chatId := message.Chat.ID
	replyText := ""

	if repo, _ := model.FindRepoByFullName(message.CommandArguments()); repo != nil {
		suber := new(model.Subscriber)
		suber.ChatId = chatId
		suber.UserName = message.Chat.UserName
		rList, err := model.FindReposBySubscriber(suber)
		if err == nil {
			for _, r := range rList {
				if r.FullName == repo.FullName {
					replyText = "无需重复订阅"
					sendText(chatId, replyText)
					return
				}
			}
		}
		model.AddRepoWithSubscriber(repo, suber)
		replyText = "订阅成功"
	} else {
		replyText = "订阅失败，请联系仓库管理员设置该仓库的 webhook"
	}

	sendText(chatId, replyText)
}

// 处理 unsub 命令
func unsub(update *tgbotapi.Update) {
	if !verifyId(update) {
		return
	}
	message := update.Message
	chatId := message.Chat.ID
	replyText := ""

	suber := new(model.Subscriber)
	suber.ChatId = chatId
	suber.UserName = message.Chat.UserName

	repo := new(model.Repository)
	repo.FullName = message.CommandArguments()
	rList, err := model.FindReposBySubscriber(suber)
	if err == nil {
		for _, r := range rList {
			if r.FullName == repo.FullName {
				model.DeleteRepoWithSubscriber(repo, suber)
				replyText = "取消订阅成功"
				sendText(chatId, replyText)
				return
			}
		}
		replyText = "该仓库未订阅，无需取消"
	} else {
		replyText = "你还未订阅任何仓库"
	}

	sendText(chatId, replyText)
}

// 处理 list 命令
func listCommand(update *tgbotapi.Update)  {
	if !verifyId(update) {
		return
	}
	message := update.Message
	chatId := message.Chat.ID
	replyText := ""

	suber := new(model.Subscriber)
	suber.ChatId = chatId
	suber.UserName = message.Chat.UserName

	rList, _ := model.FindReposBySubscriber(suber)
	if rList != nil && len(rList) == 0 {
		replyText = "还没有任何订阅"
	} else {
		replyText = fmt.Sprintf("订阅列表: \n")
		for _, r := range rList {
			replyText = fmt.Sprintf("%v\n[%v](https://github.com/%v)", replyText, r.FullName, r.FullName)
		}
	}

	sendText(chatId, replyText)
}

// 处理 all 命令
func allCommand(update *tgbotapi.Update) {
	if !verifyId(update) {
		return
	}
	message := update.Message
	chatId := message.Chat.ID
	replyText := ""

	rList, _ := model.FindAllRepos()
	if rList != nil && len(rList) == 0 {
		replyText = "还没有任何订阅"
	} else {
		replyText = fmt.Sprintf("可订阅仓库: \n")
		for _, r := range rList {
			replyText = fmt.Sprintf("%v\n[%v](https://github.com/%v)", replyText, r.FullName, r.FullName)
		}
	}

	sendText(chatId, replyText)
}

// 验证身份
func verifyId(update *tgbotapi.Update) bool {
	replyText := ""
	chatId := update.Message.Chat.ID
	if r, _ := model.FindSubscriberByChatID(chatId); r == nil {
		replyText = "请先通过鉴权"
		sendText(chatId, replyText)
		return false
	}
	return true
}

// 发送文本消息
func sendText(chatId int64, replyText string)  {
	replyMessage := tgbotapi.NewMessage(chatId, replyText)
	replyMessage.ParseMode = tgbotapi.ModeMarkdown
	replyMessage.DisableWebPagePreview = true
	_, err := tgbot.GetBot().Send(replyMessage)
	if err != nil {
		log.Println(err)
	}
}