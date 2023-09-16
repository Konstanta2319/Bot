package main

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Message struct {
	path     string
	caption  string
	Keyboard tgbotapi.ReplyKeyboardMarkup
}

var Messages = map[string]Message{
	"детская одежда": {"3.jpg", "А теперь интеесно на сколько все грамотно работает", kb1},
	"мужская одежда": {"4.jpg", "А теперь интеесно на сколько все грамотно работает", kb2},
	"женская одежда": {"5.jpg", "А теперь интеесно на сколько все грамотно работает", kb3},
	"обувь":          {"6.jpg", "А теперь интеесно на сколько все грамотно работает", kb4},
	"головные уборы": {"7.jpg", "А теперь интеесно на сколько все грамотно работает", kb5},
	"аксессуары":     {"8.jpg", "А теперь интеесно на сколько все грамотно работает", kb6},
}

func (bot *Bot) sendMessage(chatId int64, message Message) {
	msg := tgbotapi.NewPhoto(chatId, tgbotapi.FilePath("./photo/"+message.path))
	msg.Caption = message.caption
	msg.ReplyMarkup = message.Keyboard
	bot.Send(msg)
}
