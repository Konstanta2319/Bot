package main

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Photo struct {
	path     string
	caption  string
	keyboard tgbotapi.InlineKeyboardMarkup
}

var Photos = map[string]Photo{
	"👔рубашки🚼":  {"1.jpg", "Детские рубашки по выгодным ценам...", numericKeyboard11},
	"👕футболки🚼": {"2.jpg", "Детские футболки по выгодным ценам...", numericKeyboard12},
}

func (bot *Bot) sendPhoto(chatId int64, photo Photo) {
	msg := tgbotapi.NewPhoto(chatId, tgbotapi.FilePath("./photo/"+photo.path))
	msg.Caption = photo.caption
	msg.ReplyMarkup = photo.keyboard
	bot.Send(msg)
}
