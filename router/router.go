package router

import (
	"github-notify-bot/controller"
	"github.com/gin-gonic/gin"
)

// InitRouter :  初始化路由
func InitRouter() *gin.Engine {

	r := gin.Default()

	r.POST("/github-webhook", controller.ParseGithubMessage)
	r.POST("/tgbot-webhook", controller.AuthTelegramUser)

	return r
}
