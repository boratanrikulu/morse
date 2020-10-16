package main

import (
	"log"
	"os"
	"strings"

	t "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"

	m "github.com/boratanrikulu/morse/lib"
)

func init() {
	godotenv.Load(".env")
}

func main() {
	mm := m.NewMorse()

	bot, err := t.NewBotAPI(os.Getenv("TELEGRAM_API_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := t.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		switch update.Message.Command() {
		case "encode":
			text := ""
			result, err := mm.Encode(strings.NewReader(update.Message.CommandArguments()))
			text = result
			if err != nil {
				text = err.Error()
			}

			msg := t.NewMessage(update.Message.Chat.ID, text)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		case "decode":
			text := ""
			result, err := mm.Decode(strings.NewReader(update.Message.CommandArguments()))
			text = result
			if err != nil {
				text = err.Error()
			}

			msg := t.NewMessage(update.Message.Chat.ID, text)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		case "help", "start":
			msg := t.NewMessage(update.Message.Chat.ID, `You can use "/encode <TEXT>" or "/decode <TEXT>" `)

			bot.Send(msg)
		default:
			msg := t.NewMessage(update.Message.Chat.ID, `Command is not found. Check "/help".`)

			bot.Send(msg)
		}

	}
}
