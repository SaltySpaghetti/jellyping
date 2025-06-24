package routes

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"saltyspaghetti.dev/jellyping/internal/bot"
	"saltyspaghetti.dev/jellyping/models"
	"saltyspaghetti.dev/jellyping/services"
)

type UserRouter struct {
	userService *services.UserService
	bot         *bot.Bot
}

func NewUserRouter(userService *services.UserService, bot *bot.Bot) *UserRouter {
	return &UserRouter{userService: userService, bot: bot}
}

func (userRouter *UserRouter) Notify(c *gin.Context) {
	var body models.JellyseerPayload
	if err := c.ShouldBindJSON(&body); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(400, gin.H{
			"error": "Invalid request payload",
		})
		return
	}

	user, err := userRouter.userService.GetByUsername(body.Request.RequestedByUsername)
	if err != nil {
		log.Printf("User not found, skipping notification")
		c.JSON(404, gin.H{
			"error": "User not found",
		})
		return
	}

	if !user.ChatID.Valid {
		log.Printf("User %s does not have a valid chat ID, skipping", user.Username)
		c.JSON(400, gin.H{
			"error": "User does not have a valid chat ID",
		})
		return
	}

	photo := tgbotapi.NewPhoto(user.ChatID.Int64, tgbotapi.FileURL(body.Image))
	photo.Caption = fmt.Sprintf("%s\n%s", body.Subject, body.Message)

	_, err = userRouter.bot.Instance.Send(photo)
	if err != nil {
		log.Fatal(err)
		c.JSON(500, gin.H{
			"error": "Failed to send message",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Notification sent successfully",
		"user":    user,
	})
}
