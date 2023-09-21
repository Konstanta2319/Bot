package main

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Message struct {
	path     string
	caption  string
	Keyboard tgbotapi.ReplyKeyboardMarkup
}

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
		tgbotapi.NewKeyboardButton("Аксессуары"),
		tgbotapi.NewKeyboardButton("Обувь"),
	),
)

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
