package main

import (
	"log"
	"time"

	mgo "gopkg.in/mgo.v2"
	tb "gopkg.in/tucnak/telebot.v2"
)

var session *mgo.Session

func main() {
	b, err := tb.NewBot(tb.Settings{
		Token: "586866387:AAHmxTxHOUxZyjhauJ3yxedpPTWUpNxLUQE", // t.me/waves_vote_bot  для Никиты
		// Token:  "595106358:AAFyY_w1SNHReDF2j9eQQjhNHBIhElDU_QY", // t.me/test_waves_vote_bot для Кирилла
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	// Кнопки и меню
	mainData := tb.ReplyButton{Text: "💳 Мой кабинет"}
	votingData := tb.ReplyButton{Text: "Текущие голосования"}
	mainMenu := [][]tb.ReplyButton{{mainData}, {votingData}}

	viewNew := tb.InlineButton{Unique: "viewnew", Text: "Текущие голосования"}
	viewRes := tb.InlineButton{Unique: "viewres", Text: "История моих голосований"}
	viewMenu := [][]tb.InlineButton{{viewNew}, {viewRes}}
	// Кнопки и меню

	// Обработчики на главное меню
	b.Handle("/start", func(m *tb.Message) {
		// var userID = strconv.Itoa(m.Sender.ID)
		// var name = string(m.Sender.Username)
		// if userlogic.Auth(session, userID) {
		// 	userlogic.Register(session, userID, name)
		// 	var msg = "Вы зарегистрированы в системе!\n\n"
		// 	msg += "Ваш *Seed:* "
		// 	msg += seed
		// 	msg += "\n\n"
		// 	msg += "Ваш *Address:* "
		// 	msg += address
		// 	b.Send(m.Sender, msg, &tb.SendOptions{DisableWebPagePreview: true, ParseMode: "Markdown"})
		// }
		b.Send(m.Sender, "Главное меню", &tb.ReplyMarkup{ResizeReplyKeyboard: true, ReplyKeyboard: mainMenu})
	})
	b.Handle(&mainData, func(m *tb.Message) {
		var msg = "Ваш публичный адрес: "
		// msg += addressPub
		msg += "\n\nВаш seed: "
		// msg += seed
		msg += "\n\nВаш баланс: "
		// msg += balance
		msg += " (РУБ)"
		msg += "\n\nТут вы можете найти открытые голосования и посмотреть результаты прошлых"

		b.Send(m.Sender, msg, &tb.ReplyMarkup{InlineKeyboard: viewMenu})
	})
	b.Start()
	// Обработчики на главное меню
}
