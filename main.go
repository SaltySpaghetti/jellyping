package main

import (
	"context"
	"embed"
	"log"
	"os"

	"github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
	"github.com/pressly/goose/v3"
	"saltyspaghetti.dev/jellyping/internal"
	"saltyspaghetti.dev/jellyping/internal/bot"
	"saltyspaghetti.dev/jellyping/models"
	"saltyspaghetti.dev/jellyping/services"
	"saltyspaghetti.dev/jellyping/utils"
)

//go:embed internal/db/migrations/*.sql
var embedMigrations embed.FS

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	databaseUrl := os.Getenv("DATABASE_URL")
	print("Connecting to database at: ", databaseUrl)
	conn, err := utils.ConnectDatabase(databaseUrl)
	if err != nil {
		panic(err)
	}

	// Goose migration setup
	goose.SetBaseFS(embedMigrations)
	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	db := stdlib.OpenDB(*conn.Config())
	if err := goose.Up(db, "internal/db/migrations"); err != nil {
		panic(err)
	}

	userService := services.NewUserService(context.Background(), conn)

	token := os.Getenv("BOT_TOKEN")
	botInstance, err := bot.NewBot(token, userService, true)
	if err != nil {
		log.Fatal(err)
	}

	app := internal.NewApp(
		models.NewConfig(),
		conn,
		botInstance,
		userService,
	)

	app.Bot.SetupAndRun()

	app.ImportUsers()
	app.Run()
}
