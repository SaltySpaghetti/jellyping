package routes

import (
	"log"

	"saltyspaghetti.dev/jellyping/internal/bot"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"saltyspaghetti.dev/jellyping/services"
)

func NewRouter(conn *pgx.Conn, bot *bot.Bot, userService *services.UserService) *gin.Engine {
	router := gin.Default()
	err := router.SetTrustedProxies(nil)
	if err != nil {
		log.Fatal(err)
	}

	group := router.Group("/api/v1")

	group.GET("/health", func(c *gin.Context) {
		c.JSON(200, nil)
	})

	userRouter := NewUserRouter(userService, bot)
	group.POST("/notify", userRouter.Notify)

	return router
}
