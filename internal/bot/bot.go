package bot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"saltyspaghetti.dev/jellyping/services"
)

type Bot struct {
	Instance    *tgbotapi.BotAPI
	userService *services.UserService
}

func NewBot(token string, userService *services.UserService, debug ...bool) (*Bot, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	if debug != nil {
		bot.Debug = debug[0]
	} else {
		bot.Debug = false
	}

	return &Bot{bot, userService}, nil
}

func (bot *Bot) SetupAndRun() {
	log.Printf("Authorized on account %s", bot.Instance.Self.UserName)
	config := tgbotapi.NewUpdate(0)
	config.Timeout = 30

	go bot.run(config)
}

func (bot *Bot) run(updateConfig tgbotapi.UpdateConfig) {
	updates := bot.Instance.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil || !update.Message.IsCommand() {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		switch update.Message.Command() {
		case "username":
			_, err := bot.userService.SetChatId(
				update.Message.CommandArguments(),
				update.Message.Chat.ID,
			)

			if err != nil {
				log.Printf("Error setting username: %v", err)
				switch err.Error() {
				case "user not found":
					msg.Text = "Username non trovato. Assicurati che il tuo username sia corretto."
				default:
					msg.Text = "Errore durante l'impostazione dell'username."
				}
			} else {
				msg.Text = "Username impostato!"
			}

		default:
			msg.Text = `
				Lista comandi:
				- /username <tuo username>
			`
		}

		if _, err := bot.Instance.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
