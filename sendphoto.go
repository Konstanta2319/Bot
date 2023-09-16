package main

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Photo struct {
	path    string
	caption string
	link    string
}

var Photos = map[string]Photo{
	"👔Рубашки🚼":    {"1.jpg", "Детские рубашки по выгодным ценам...", "https://www.wildberries.ru/catalog/detyam/odezhda/dlya-malchikov/rubashki"},
	"👕футболки🚼":   {"2.jpg", "Детские футболки по выгодным ценам...", "https://www.ozon.ru/category/futbolki-i-polo-dlya-malchikov/"},
	"🦺майки🚼":      {"9.jpg", "Детские майки по выгодным ценам...", "https://www.ozon.ru/category/detskie-mayki/"},
	"🩳шорты🚼":      {"10.jpg", "Детские шорты по выгодным ценам...", "https://www.ozon.ru/category/shorty-dlya-malchikov/"},
	"👖джинсы🚼":     {"11.jpg", "Детские джинсы по выгодным ценам...", "https://www.ozon.ru/category/shorty-dlya-malchikov/"},
	"👖штаны🚼":      {"12.jpg", "Детские штаны по выгодным ценам...", "https://www.ozon.ru/category/shorty-dlya-malchikov/"},
	"👘платья🚼":     {"13.jpg", "Детские платья по выгодным ценам...", "https://www.ozon.ru/category/shorty-dlya-malchikov/"},
	"👗юбки🚼":       {"14.jpg", "Детские юбки по выгодным ценам...", "https://www.ozon.ru/category/shorty-dlya-malchikov/"},
	"🥼блузки🚼":     {"15.jpeg", "Детские блузки по выгодным ценам...", "https://www.ozon.ru/category/shorty-dlya-malchikov/"},
	"👔рубашки🚹":    {"16.jpg", "Мужские рубашки по выгодным ценам...", "https://www.ozon.ru/category/shorty-dlya-malchikov/"},
	"👕футболки🚹":   {"17.jpg", "Футболки по выгодным ценам...", "https://www.ozon.ru/category/shorty-dlya-malchikov/"},
	"🦺майки🚹":      {"18.jpg", "Майки по выгодным ценам...", "https://www.ozon.ru/category/shorty-dlya-malchikov/"},
	"🩳шорты🚹":      {"19.jpg", "Шорты по выгодным ценам...", "https://www.ozon.ru/category/shorty-dlya-malchikov/"},
	"👖джинсы🚹":     {"20.jpg", "Джинсы по выгодным ценам...", "https://www.ozon.ru/category/shorty-dlya-malchikov/"},
	"👖штаны🚹":      {"21.jpg", "Штаны по выгодным ценам...", "https://www.ozon.ru/category/shorty-dlya-malchikov/"},
	"👘платья🚺":     {"22.jpg", "Платья по выгодным ценам...", "https://www.ozon.ru/category/shorty-dlya-malchikov/"},
	"👗юбки🚺":       {"23.jpg", "Юбки по выгодным ценам...", "https://www.ozon.ru/category/shorty-dlya-malchikov/"},
	"🥼блузки🚺":     {"24.jpg", "Блузки по выгодным ценам...", "https://www.ozon.ru/category/shorty-dlya-malchikov/"},
	"🩳шорты🚺":      {"25.jpg", "Женские шорты по выгодным ценам...", "https://www.ozon.ru/category/shorty-dlya-malchikov/"},
	"👖джинсы🚺":     {"26.jpg", "Женские джинсы по выгодным ценам...", "https://www.ozon.ru/category/shorty-dlya-malchikov/"},
	"👖штаны🚺":      {"27.jpg", "Женские штаны по выгодным ценам...", "https://www.ozon.ru/category/shorty-dlya-malchikov/"},
	"👟кросовки👟":   {"28.jpg", "Кросовки по выгодным ценам...", "https://www.ozon.ru/category/shorty-dlya-malchikov/"},
	"👞туфли👞":      {"29.jpg", "Туфли по выгодным ценам...", "https://www.ozon.ru/category/shorty-dlya-malchikov/"},
	"🥾ботинки🥾":    {"30.jpg", "Ботинки по выгодным ценам...", "https://www.ozon.ru/category/shorty-dlya-malchikov/"},
	"🩴тапочки🩴":    {"31.jpg", "Тапочки по выгодным ценам...", "https://www.ozon.ru/category/shorty-dlya-malchikov/"},
	"👢зимняя👢":     {"32.jpg", "Зимняя обувь по выгодным ценам...", "https://www.ozon.ru/category/shorty-dlya-malchikov/"},
	"👟всесезоняя👠": {"33.jpg", "Всесезоняя обувь по выгодным ценам...", "https://www.ozon.ru/category/shorty-dlya-malchikov/"},
	"🧢кепки🧢":      {"34.jpg", "Кепки по выгодным ценам...", "https://www.ozon.ru/category/shorty-dlya-malchikov/"},
	"👒панамы👒":     {"35.jpg", "Панамы по выгодным ценам...", "https://www.ozon.ru/category/shorty-dlya-malchikov/"},
	"👑шапки👑":      {"36.jpg", "Шапки по выгодным ценам...", "https://www.ozon.ru/category/shorty-dlya-malchikov/"},
	"✨колье✨":      {"37.jpg", "Женские украшения по выгодным ценам...", "https://www.ozon.ru/category/shorty-dlya-malchikov/"},
	"💍кольца💍":     {"38.jpg", "Женские украшения по выгодным ценам...", "https://www.ozon.ru/category/shorty-dlya-malchikov/"},
	"✨серьги✨":     {"39.jpg", "Женские украшения по выгодным ценам...", "https://www.ozon.ru/category/shorty-dlya-malchikov/"},
}

func (bot *Bot) sendPhoto(chatId int64, photo Photo) {
	msg := tgbotapi.NewPhoto(chatId, tgbotapi.FilePath("./photo/"+photo.path))
	msg.Caption = photo.caption
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL("Заказать вы можете тут 👉", photo.link),
		),
	)

	bot.Send(msg)
}
