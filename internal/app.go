package internal

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
	"saltyspaghetti.dev/jellyping/internal/bot"
	"saltyspaghetti.dev/jellyping/models"
	"saltyspaghetti.dev/jellyping/routes"
	"saltyspaghetti.dev/jellyping/services"
)

type App struct {
	config      *models.Config
	Conn        *pgx.Conn
	Bot         *bot.Bot
	UserService *services.UserService
}

func (app *App) Run() {
	log.Printf("Starting server on port %s", app.config.Port)

	router := routes.NewRouter(app.Conn, app.Bot, app.UserService)
	err := router.Run(":" + app.config.Port)
	if err != nil {
		log.Fatal(err)
	}
}

func (app *App) ImportUsers() {
	jellyfinUrl := os.Getenv("JELLYFIN_URL")
	if jellyfinUrl == "" {
		jellyfinUrl = "http://localhost:8096"
	}

	jellyfinApiKey := os.Getenv("JELLYFIN_API_KEY")
	if jellyfinApiKey == "" {
		log.Println("JELLYFIN_API_KEY is not set. Skipping user import.")
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", jellyfinUrl+"/Users", nil)
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("MediaBrowser Token=%s", jellyfinApiKey))
	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to make request: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatalf("Failed to fetch users: %s", res.Status)
	}

	var users []models.JellyfinUser
	if err := json.NewDecoder(res.Body).Decode(&users); err != nil {
		log.Fatalf("Failed to decode response: %v", err)
	}

	log.Printf("Found %d users in Jellyfin", len(users))
	
	importedUsers := 0
	for _, user := range users {
		foundUser, err := app.UserService.UserExists(user.Name)
		if err != nil {
			log.Printf("Error checking user %s: %v", user.Name, err)
			continue
		}

		if foundUser {
			log.Printf("User %s already exists, skipping", user.Name)
			continue
		}

		app.UserService.CreateUser(user.Name)
		importedUsers++
	}

	log.Printf("Imported %d users from Jellyfin", importedUsers)
}

func NewApp(
	config *models.Config,
	conn *pgx.Conn,
	bot *bot.Bot,
	userService *services.UserService,
) *App {
	return &App{
		config:      config,
		Conn:        conn,
		Bot:         bot,
		UserService: userService,
	}
}
