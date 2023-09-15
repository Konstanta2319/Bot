package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strings"
)

var numericKeyboard1 = tgbotapi.NewReplyKeyboard(
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

var numericKeyboard2 = tgbotapi.NewReplyKeyboard(
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

var numericKeyboard3 = tgbotapi.NewReplyKeyboard(
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

var numericKeyboard4 = tgbotapi.NewReplyKeyboard(
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

var numericKeyboard5 = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("🧢Кепки🧢"),
		tgbotapi.NewKeyboardButton("👒Панамы👒"),
		tgbotapi.NewKeyboardButton("👑Шапки👑"),
	), tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("🔙назад"),
	),
)

var numericKeyboard6 = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("✨Колье✨"),
		tgbotapi.NewKeyboardButton("💍Кольца💍"),
		tgbotapi.NewKeyboardButton("✨Серьги✨"),
	), tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("🔙назад"),
	),
)

var numericKeyboard11 = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL("Заказать вы можете тут 👉", "https://www.wildberries.ru/catalog/detyam/odezhda/dlya-malchikov/rubashki"),
	),
)
var numericKeyboard12 = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL("Заказать вы можете тут 👉", "https://www.ozon.ru/category/futbolki-i-polo-dlya-malchikov/"),
	),
)
var numericKeyboard13 = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL("Заказать вы можете тут 👉", "https://www.ozon.ru/category/detskie-mayki/"),
	),
)
var numericKeyboard14 = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL("Заказать вы можете тут 👉", "https://www.ozon.ru/category/shorty-dlya-malchikov/"),
	),
)
var numericKeyboard15 = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL("Заказать вы можете тут 👉", "https://www.ozon.ru/category/shorty-dlya-malchikov/"),
	),
)
var numericKeyboard16 = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL("Заказать вы можете тут 👉", "https://www.ozon.ru/category/shorty-dlya-malchikov/"),
	),
)
var numericKeyboard17 = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL("Заказать вы можете тут 👉", "https://www.ozon.ru/category/shorty-dlya-malchikov/"),
	),
)
var numericKeyboard18 = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL("Заказать вы можете тут 👉", "https://www.ozon.ru/category/shorty-dlya-malchikov/"),
	),
)
var numericKeyboard19 = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL("Заказать вы можете тут 👉", "https://www.ozon.ru/category/shorty-dlya-malchikov/"),
	),
)
var numericKeyboard20 = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL("Заказать вы можете тут 👉", "https://www.ozon.ru/category/shorty-dlya-malchikov/"),
	),
)
var numericKeyboard21 = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL("Заказать вы можете тут 👉", "https://www.ozon.ru/category/shorty-dlya-malchikov/"),
	),
)

var numericKeyboard = tgbotapi.NewReplyKeyboard(
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
	tgbot, err := tgbotapi.NewBotAPI("5570771711:AAHhr-jw60pmQUXeWaScTtqKe8XTG8JL3_Q")
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
			} else if text == "привет" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ну привет! 😀")
				msg.ReplyMarkup = numericKeyboard
				bot.Send(msg)
			} else if text == "как дела?" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Всё отлично =) У тебя как?")
				bot.Send(msg)
			} else if text == "что ты можешь?" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "особо пока ничего нового")
				bot.Send(msg)
			} else if text == "🦺майки🚼" {
				photo := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FilePath("./photo/9.jpg"))
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Детские майки по выгодным ценам...")
				msg.ReplyMarkup = numericKeyboard13
				bot.Send(photo)
				bot.Send(msg)
			} else if text == "🩳шорты🚼" {
				photo := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FilePath("./photo/10.jpg"))
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Детские шорты по выгодным ценам...")
				msg.ReplyMarkup = numericKeyboard14
				bot.Send(photo)
				bot.Send(msg)
			} else if text == "👖джинсы🚼" {
				photo := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FilePath("./photo/11.jpg"))
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Детские джинсы по выгодным ценам...")
				msg.ReplyMarkup = numericKeyboard15
				bot.Send(photo)
				bot.Send(msg)
			} else if text == "👖штаны🚼" {
				photo := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FilePath("./photo/12.jpg"))
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Детские штаны по выгодным ценам...")
				msg.ReplyMarkup = numericKeyboard16
				bot.Send(photo)
				bot.Send(msg)
			} else if text == "👘платья🚼" {
				photo := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FilePath("./photo/13.jpg"))
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Детские платья по выгодным ценам...")
				msg.ReplyMarkup = numericKeyboard17
				bot.Send(photo)
				bot.Send(msg)
			} else if text == "👗юбки🚼" {
				photo := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FilePath("./photo/14.jpg"))
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Детские юбки по выгодным ценам...")
				msg.ReplyMarkup = numericKeyboard18
				bot.Send(photo)
				bot.Send(msg)
			} else if text == "🥼блузки🚼" {
				photo := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FilePath("./photo/15.jpeg"))
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Детские блузки по выгодным ценам...")
				msg.ReplyMarkup = numericKeyboard18
				bot.Send(photo)
				bot.Send(msg)
			} else if text == "детская одежда" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "когда нибудь, но не сегодня)")
				photo := tgbotapi.NewPhoto(update.Message.From.ID, tgbotapi.FilePath("./photo/3.jpg"))
				msg.ReplyMarkup = numericKeyboard1
				bot.Send(photo)
				bot.Send(msg)
			} else if text == "мужская одежда" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Бывало и по лучше!")
				photo := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FilePath("./photo/4.jpg"))
				msg.ReplyMarkup = numericKeyboard2
				bot.Send(photo)
				bot.Send(msg)
			} else if text == "женская одежда" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ну пока учусь")
				photo := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FilePath("./photo/5.jpg"))
				msg.ReplyMarkup = numericKeyboard3
				bot.Send(photo)
				bot.Send(msg)
			} else if text == "обувь" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ну с кем не бывает")
				photo := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FilePath("./photo/6.jpg"))
				msg.ReplyMarkup = numericKeyboard4
				bot.Send(photo)
				bot.Send(msg)
			} else if text == "головные уборы" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "зато сам, ну почти XD")
				photo := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FilePath("./photo/7.jpg"))
				msg.ReplyMarkup = numericKeyboard5
				bot.Send(photo)
				bot.Send(msg)
			} else if text == "аксесуары" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "по факту, я устал")
				photo := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FilePath("./photo/8.jpg"))
				msg.ReplyMarkup = numericKeyboard6
				bot.Send(photo)
				bot.Send(msg)
			} else if text == "👔рубашки🚹" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Мужские рубашки по выгодным ценам...")
				photo := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FilePath("./photo/16.jpg"))
				msg.ReplyMarkup = numericKeyboard11
				bot.Send(photo)
				bot.Send(msg)
			} else if text == "👕футболки🚹" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Футболки по выгодным ценам...")
				photo := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FilePath("./photo/17.jpg"))
				msg.ReplyMarkup = numericKeyboard11
				bot.Send(photo)
				bot.Send(msg)
			} else if text == "🦺майки🚹" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Майки по выгодным ценам...")
				photo := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FilePath("./photo/18.jpg"))
				msg.ReplyMarkup = numericKeyboard11
				bot.Send(photo)
				bot.Send(msg)
			} else if text == "🩳шорты🚹" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Шорты по выгодным ценам...")
				photo := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FilePath("./photo/19.jpg"))
				msg.ReplyMarkup = numericKeyboard11
				bot.Send(photo)
				bot.Send(msg)
			} else if text == "👖джинсы🚹" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Джинсы по выгодным ценам...")
				photo := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FilePath("./photo/20.jpg"))
				msg.ReplyMarkup = numericKeyboard11
				bot.Send(photo)
				bot.Send(msg)
			} else if text == "👖штаны🚹" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Штаны по выгодным ценам...")
				photo := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FilePath("./photo/21.jpg"))
				msg.ReplyMarkup = numericKeyboard11
				bot.Send(photo)
				bot.Send(msg)
			} else if text == "👘платья🚺" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Платья по выгодным ценам...")
				photo := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FilePath("./photo/22.jpg"))
				msg.ReplyMarkup = numericKeyboard11
				bot.Send(photo)
				bot.Send(msg)
			} else if text == "👗юбки🚺" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Юбки по выгодным ценам...")
				photo := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FilePath("./photo/23.jpg"))
				msg.ReplyMarkup = numericKeyboard11
				bot.Send(photo)
				bot.Send(msg)
			} else if text == "🥼блузки🚺" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Блузки по выгодным ценам...")
				photo := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FilePath("./photo/24.jpg"))
				msg.ReplyMarkup = numericKeyboard11
				bot.Send(photo)
				bot.Send(msg)
			} else if text == "🩳шорты🚺" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Женские шорты по выгодным ценам...")
				photo := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FilePath("./photo/25.jpg"))
				msg.ReplyMarkup = numericKeyboard11
				bot.Send(photo)
				bot.Send(msg)
			} else if text == "👖джинсы🚺" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Женские джинсы по выгодным ценам...")
				photo := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FilePath("./photo/26.jpg"))
				msg.ReplyMarkup = numericKeyboard11
				bot.Send(photo)
				bot.Send(msg)
			} else if text == "👖штаны🚺" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Женские штаны по выгодным ценам...")
				photo := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FilePath("./photo/27.jpg"))
				msg.ReplyMarkup = numericKeyboard11
				bot.Send(photo)
				bot.Send(msg)
			} else if text == "👟кросовки👟" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Кросовки по выгодным ценам...")
				photo := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FilePath("./photo/28.jpg"))
				msg.ReplyMarkup = numericKeyboard11
				bot.Send(photo)
				bot.Send(msg)
			} else if text == "👞туфли👞" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Туфли по выгодным ценам...")
				photo := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FilePath("./photo/29.jpg"))
				msg.ReplyMarkup = numericKeyboard11
				bot.Send(photo)
				bot.Send(msg)
			} else if text == "🥾ботинки🥾" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ботинки по выгодным ценам...")
				photo := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FilePath("./photo/30.jpg"))
				msg.ReplyMarkup = numericKeyboard11
				bot.Send(photo)
				bot.Send(msg)
			} else if text == "🩴тапочки🩴" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Тапочки по выгодным ценам...")
				photo := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FilePath("./photo/31.jpg"))
				msg.ReplyMarkup = numericKeyboard11
				bot.Send(photo)
				bot.Send(msg)
			} else if text == "👢зимняя👢" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Зимняя обувь по выгодным ценам...")
				photo := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FilePath("./photo/32.jpg"))
				msg.ReplyMarkup = numericKeyboard11
				bot.Send(photo)
				bot.Send(msg)
			} else if text == "👟всесезоняя👠" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Всесезоняя обувь по выгодным ценам...")
				photo := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FilePath("./photo/33.jpg"))
				msg.ReplyMarkup = numericKeyboard11
				bot.Send(photo)
				bot.Send(msg)
			} else if text == "🧢кепки🧢" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Кепки по выгодным ценам...")
				photo := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FilePath("./photo/34.jpg"))
				msg.ReplyMarkup = numericKeyboard11
				bot.Send(photo)
				bot.Send(msg)
			} else if text == "👒панамы👒" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Панамы по выгодным ценам...")
				photo := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FilePath("./photo/35.jpg"))
				msg.ReplyMarkup = numericKeyboard11
				bot.Send(photo)
				bot.Send(msg)
			} else if text == "👑шапки👑" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Шапки по выгодным ценам...")
				photo := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FilePath("./photo/36.jpg"))
				msg.ReplyMarkup = numericKeyboard11
				bot.Send(photo)
				bot.Send(msg)
			} else if text == "✨колье✨" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Женские украшения по выгодным ценам...")
				photo := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FilePath("./photo/37.jpg"))
				msg.ReplyMarkup = numericKeyboard11
				bot.Send(photo)
				bot.Send(msg)
			} else if text == "💍кольца💍" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Женские украшения по выгодным ценам...")
				photo := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FilePath("./photo/38.jpg"))
				msg.ReplyMarkup = numericKeyboard11
				bot.Send(photo)
				bot.Send(msg)
			} else if text == "✨серьги✨" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Женские украшения по выгодным ценам...")
				photo := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FilePath("./photo/39.jpg"))
				msg.ReplyMarkup = numericKeyboard11
				bot.Send(photo)
				bot.Send(msg)
			} else if text == "🔙назад" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "🔙назад")
				msg.ReplyMarkup = numericKeyboard
				bot.Send(msg)
			} else if text == "кто ты?" {
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
