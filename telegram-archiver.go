package main

import (
  "github.com/go-telegram-bot-api/telegram-bot-api"
  "log"
  "os"
  "fmt"
  "strconv"
  "time"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("181465813:AAGzyXKZddJkZ9F1kv35qgZIG6wX9lPwXq8")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
    log.Printf("chat: 0b%b,0x%x\n", update.Message.Chat.ID, update.Message.Chat.ID)

    f, err := os.OpenFile(strconv.FormatInt(update.Message.Chat.ID, 10) + ".txt", os.O_APPEND|os.O_CREATE, 0666)
    if err != nil {
      log.Panic(err)
    }
    f.WriteString(fmt.Sprintf("%s [%s] %s\n", time.Unix(int64(update.Message.Date), 0), update.Message.From.UserName, update.Message.Text))
    f.Close()


	}
}
