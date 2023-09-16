package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strings"
)

var kb1 = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("👔Рубашки🚼"),
		tgbotapi.NewKeyboardButton("👕Футболки🚼"),
		tgbotapi.NewKeyboardButton("🦺Майки🚼"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("🩳Шорты🚼"),
		tgbotapi.NewKeyboardButton("👖Джинсы🚼"),
		tgbotapi.NewKeyboardButton("👖Штаны🚼"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("👘Платья🚼"),
		tgbotapi.NewKeyboardButton("👗Юбки🚼"),
		tgbotapi.NewKeyboardButton("🥼Блузки🚼"),
	), tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("🔙назад"),
	),
)

var kb2 = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("👔Рубашки🚹"),
		tgbotapi.NewKeyboardButton("👕Футболки🚹"),
		tgbotapi.NewKeyboardButton("🦺Майки🚹"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("🩳Шорты🚹"),
		tgbotapi.NewKeyboardButton("👖Джинсы🚹"),
		tgbotapi.NewKeyboardButton("👖Штаны🚹"),
	), tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("🔙назад"),
	),
)

var kb3 = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("👘Платья🚺"),
		tgbotapi.NewKeyboardButton("👗Юбки🚺"),
		tgbotapi.NewKeyboardButton("🥼Блузки🚺"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("🩳Шорты🚺"),
		tgbotapi.NewKeyboardButton("👖Джинсы🚺"),
		tgbotapi.NewKeyboardButton("👖Штаны🚺"),
	), tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("🔙назад"),
	),
)

var kb4 = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("👟Кросовки👟"),
		tgbotapi.NewKeyboardButton("👞Туфли👞"),
		tgbotapi.NewKeyboardButton("🥾Ботинки🥾"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("🩴Тапочки🩴"),
		tgbotapi.NewKeyboardButton("👢Зимняя👢"),
		tgbotapi.NewKeyboardButton("👟Всесезоняя👠"),
	), tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("🔙назад"),
	),
)

var kb5 = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("🧢Кепки🧢"),
		tgbotapi.NewKeyboardButton("👒Панамы👒"),
		tgbotapi.NewKeyboardButton("👑Шапки👑"),
	), tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("🔙назад"),
	),
)

var kb6 = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("✨Колье✨"),
		tgbotapi.NewKeyboardButton("💍Кольца💍"),
		tgbotapi.NewKeyboardButton("✨Серьги✨"),
	), tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("🔙назад"),
	),
)

var kb = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Детская одежда"),
		tgbotapi.NewKeyboardButton("Мужская одежда"),
		tgbotapi.NewKeyboardButton("Женская одежда"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Головные уборы"),
		tgbotapi.NewKeyboardButton("Аксесуары"),
		tgbotapi.NewKeyboardButton("Обувь"),
		//tgbotapi.NewKeyboardButton("Список"),
	),
)

type Bot struct {
	*tgbotapi.BotAPI
}

func main() {
	tgbot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	tgbot.Debug = true
	bot := &Bot{
		BotAPI: tgbot,
	}
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			chatId := update.Message.Chat.ID

			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			text := strings.ToLower(update.Message.Text)
			photo, ok := Photos[text]
			if ok {
				bot.sendPhoto(chatId, photo)
				continue
			}

			message, ok := Messages[text]
			if ok {
				bot.sendMessage(chatId, message)
				continue
			}

			if text == "привет" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ну привет! 😀")
				msg.ReplyMarkup = kb
				bot.Send(msg)
			} else if text == "как дела?" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Всё отлично =) У тебя как?")
				bot.Send(msg)
			} else if text == "что ты можешь?" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "особо пока ничего нового")
				bot.Send(msg)
			} else if text == "мужская одежда" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Бывало и по лучше!")
				photo := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FilePath("./photo/4.jpg"))
				msg.ReplyMarkup = kb2
				bot.Send(photo)
				bot.Send(msg)
			} else if text == "женская одежда" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ну пока учусь")
				photo := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FilePath("./photo/5.jpg"))
				msg.ReplyMarkup = kb3
				bot.Send(photo)
				bot.Send(msg)
			} else if text == "обувь" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ну с кем не бывает")
				photo := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FilePath("./photo/6.jpg"))
				msg.ReplyMarkup = kb4
				bot.Send(photo)
				bot.Send(msg)
			} else if text == "головные уборы" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "зато сам, ну почти XD")
				photo := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FilePath("./photo/7.jpg"))
				msg.ReplyMarkup = kb5
				bot.Send(photo)
				bot.Send(msg)
			} else if text == "аксессуары" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "по факту, я устал")
				photo := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FilePath("./photo/8.jpg"))
				msg.ReplyMarkup = kb6
				bot.Send(photo)
				bot.Send(msg)
			} else if text == "🔙назад" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "🔙назад")
				msg.ReplyMarkup = kb
				bot.Send(msg)
			} else if text == "кто ты?" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Я автономная программа, создан для облегчения жизни вам")
				bot.Send(msg)
			} else if text == "что ты?" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Я автономная программа, создан для облегчения жизни вам")
				bot.Send(msg)
			} else {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "я не понимю, простите😭")
				bot.Send(msg)
			}

		}

		if update.CallbackQuery != nil {
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			if _, err := bot.Request(callback); err != nil {
				panic(err)
			}

			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
			if _, err := bot.Send(msg); err != nil {
				panic(err)
			}

		}

	}

}
