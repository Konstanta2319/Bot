package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"math/rand"
	"strings"
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

			answers := []string{
				"Привет",
				"Hi",
				"Приветсвую",
				"Каничива",
				"Здравствуйте",
				"Как ваши ничего?",
				"Салам",
				"Саламалейкум",
				"Namaskaar",
				"你好! ",
				"안녕 ",
				"Sain uu?",
				"どうも",
				"Sua s'dei",
				"Wai",
				"Halo",
				"Salem",
			}

			answers2 := []string{
				"Нормально,у вас как?",
				"Бывало и получше)",
				"Да в целом, все норм",
				"Сойдет",
				"Я рад что вы спросили)",
				"Да все хорошо,вы как?",
				"До этого момента все было отлично, а теперь стало супер)",
			}

			if text == "привет" {
				msg := tgbotapi.NewMessage(chatId, answers[rand.Intn(len(answers))])
				msg.ReplyMarkup = kb
				bot.Send(msg)
			} else if text == "как дела?" {
				msg := tgbotapi.NewMessage(chatId, answers2[rand.Intn(len(answers2))])
				bot.Send(msg)
			} else if text == "что ты можешь?" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "особо пока ничего нового")
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
