package main

import (
	"github.com/tirdman/tg-guess-number-golang/config"
	"github.com/tirdman/tg-guess-number-golang/utils"
	"gopkg.in/telegram-bot-api.v4"
	"log"
	"strconv"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(config.Token)
	if err != nil {
		log.Panic(err)
	}

	unknownNumer := utils.GenerateNum(4)
	var usersAttempts []int

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		text := update.Message.Text

		if !utils.IsNumber(text) || len(text) != 4 {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID,
				"Введенный текст должен содержать только 4-значное число. Повторите ввод.")
			bot.Send(msg)
			continue
		}

		usersAttempts = append(usersAttempts, update.Message.From.ID)
		usersAttempt := utils.GetAttemptsByUser(usersAttempts, update.Message.From.ID)
		answer := utils.CheckInputNumber(text, unknownNumer)

		if answer == "BBBB" {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID,
				"Поздравляем! Вы угадали число "+text+" с попытки: "+strconv.Itoa(usersAttempt)+". Сгенерировано новое число.")
			bot.Send(msg)

			allUserInQuest := utils.GetAllUserInCurrentQuest(usersAttempts)
			for _, v := range allUserInQuest {
				if v == update.Message.From.ID {
					continue
				}

				msg := tgbotapi.NewMessage(int64(v),
					"Игра завершена. Игрок "+update.Message.From.FirstName+" "+update.Message.From.LastName+
						" угадал число "+unknownNumer+" c попытки: "+strconv.Itoa(usersAttempt)+". Сгенерировано новое число.")
				bot.Send(msg)

			}

			unknownNumer = utils.GenerateNum(4)
			usersAttempts = usersAttempts[:0]
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Попытка: "+strconv.Itoa(usersAttempt)+". Не угадали: "+answer)
		bot.Send(msg)
	}
}
