package main

import (
	"github.com/tirdman/tg-guess-number-golang/config"
	"github.com/tirdman/tg-guess-number-golang/models"
	"github.com/tirdman/tg-guess-number-golang/utils"
	"gopkg.in/telegram-bot-api.v4"
	"log"
	"strconv"
	"strings"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(config.Token)
	if err != nil {
		log.Panic(err)
	}

	unknownNumer := utils.GenerateNum(4)
	var users []*models.User

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		text := strings.TrimSpace(update.Message.Text)

		if text == "/start" {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID,
				"Угадайте "+strconv.Itoa(len(unknownNumer))+"-значное число.")
			bot.Send(msg)
			continue
		}

		if !utils.IsNumber(text) || len(text) != len(unknownNumer) {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID,
				"Введенный текст должен содержать только "+strconv.Itoa(len(unknownNumer))+"-значное число. Повторите ввод.")
			bot.Send(msg)
			continue
		}

		u, err := utils.GetUser(update.Message.From.ID, users)
		if err != nil {
			u = &models.User{update.Message.From.ID, 1}
			users = append(users, u)
		} else {
			u.Attempts++
		}

		answer := utils.CheckInputNumber(text, unknownNumer)

		if answer == "BBBB" {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID,
				"Поздравляем! Вы угадали число "+text+" с попытки: "+strconv.Itoa(u.Attempts)+". Сгенерировано новое число.")
			bot.Send(msg)

			for _, v := range users {
				if v.Id == update.Message.From.ID {
					continue
				}

				msg := tgbotapi.NewMessage(int64(v.Id),
					"Игра завершена. Игрок "+update.Message.From.FirstName+" "+update.Message.From.LastName+
						" угадал число "+unknownNumer+" c попытки: "+strconv.Itoa(u.Attempts)+". Сгенерировано новое число.")
				bot.Send(msg)

			}

			unknownNumer = utils.GenerateNum(4)
			users = users[:0]
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Попытка: "+strconv.Itoa(u.Attempts)+". Не угадали: "+answer)
		bot.Send(msg)
	}
}
