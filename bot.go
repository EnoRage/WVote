package main

import (
	"log"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

var proj1 = "*Организация:* Фонд НеСкам\n\n*Адрес:* liwjbfiowbupiweubpwbep\n\n*Тема голосования:* Голосуем за то, чтобы провести новое ICO\n\n*Условия голосования:* иметь больше 100 noscum токенов\n\n*Время завершения голосования:* 9 Апреля, 00:01 (MSK)"
var proj2 = "*Организация:*  ------\n\n*Адрес:* ------\n\n*Тема голосования:* ------\n\n*Условия голосования:* ------\n\n*Время завершения голосования:* ------"
var proj3 = "*Организация:*  ---------\n\n*Адрес:* ------\n\n*Тема голосования:* ---------\n\n*Условия голосования:* ---------\n\n*Время завершения голосования:* ------"
var proj4 = "*Организация:*  ---------------\n\n*Адрес:* ---------\n\n*Тема голосования:* ------------\n\n*Условия голосования:* ------------\n\n*Время завершения голосования:* ------"
var choseproj = ""
var yesnores = ""

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

	viewRes := tb.InlineButton{Unique: "viewres", Text: "История моих голосований"}
	viewMenu := [][]tb.InlineButton{{viewRes}}

	listVote1 := tb.InlineButton{Unique: "listvote1", Text: "Страница 1"}
	listVote2 := tb.InlineButton{Unique: "listvote2", Text: "Страница 2"}
	listVote3 := tb.InlineButton{Unique: "listvote3", Text: "Страница 3"}
	listVote4 := tb.InlineButton{Unique: "listvote4", Text: "Страница 4"}

	vote1 := tb.InlineButton{Unique: "vote1", Text: "Проголосовать за 1"}
	vote2 := tb.InlineButton{Unique: "vote2", Text: "Проголосовать за 2"}
	vote3 := tb.InlineButton{Unique: "vote3", Text: "Проголосовать за 3"}
	vote4 := tb.InlineButton{Unique: "vote4", Text: "Проголосовать за 4"}

	menu := tb.InlineButton{Unique: "menu", Text: "Главное меню"}

	menuVote1 := [][]tb.InlineButton{{vote1}, {listVote2, listVote3, listVote4}, {menu}}
	menuVote2 := [][]tb.InlineButton{{vote2}, {listVote1, listVote3, listVote4}, {menu}}
	menuVote3 := [][]tb.InlineButton{{vote3}, {listVote1, listVote2, listVote4}, {menu}}
	menuVote4 := [][]tb.InlineButton{{vote4}, {listVote1, listVote2, listVote3}, {menu}}

	yes := tb.InlineButton{Unique: "yes", Text: "✅ За"}
	no := tb.InlineButton{Unique: "no", Text: "❌ Против"}
	yesno := [][]tb.InlineButton{
		{yes, no}, {menu}}

	// Кнопки и меню

	// Обработчики на главное меню
	b.Handle("/start", func(m *tb.Message) {
		b.Send(m.Sender, "Главное меню", &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{ResizeReplyKeyboard: true, ReplyKeyboard: mainMenu})
	})
	b.Handle(&mainData, func(m *tb.Message) {
		var msg = "Ваш публичный адрес: "
		// msg += addressPub
		msg += "\n\nВаш seed: "
		// msg += seed
		msg += "\n\nВаш баланс: "
		// msg += balance
		msg += " (РУБ)"
		msg += "\n\nТут вы можете посмотреть результаты прошлых голосований"
		b.Send(m.Sender, msg, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: viewMenu})
	})
	b.Handle(&viewRes, func(c *tb.Callback) {
		b.Edit(c.Message, &tb.SendOptions{ParseMode: "Markdown"}, "Тут вы можете посмотреть куда и как вы голосовали, а также узнать результат")
		b.Respond(c, &tb.CallbackResponse{})
	})
	b.Handle(&votingData, func(m *tb.Message) {
		var msg = "Страница 1:\n\n"
		msg += proj1
		b.Send(m.Sender, msg, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: menuVote1})
	})
	b.Handle(&listVote1, func(c *tb.Callback) {
		var msg = "Страница 1:\n\n"
		msg += proj1
		b.Edit(c.Message, msg, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: menuVote1})
		b.Respond(c, &tb.CallbackResponse{})
	})
	b.Handle(&listVote2, func(c *tb.Callback) {
		var msg = "Страница 2:\n\n"
		msg += proj2
		b.Edit(c.Message, msg, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: menuVote2})
		b.Respond(c, &tb.CallbackResponse{})
	})
	b.Handle(&listVote3, func(c *tb.Callback) {
		var msg = "Страница 3:\n\n"
		msg += proj3
		b.Edit(c.Message, msg, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: menuVote3})
		b.Respond(c, &tb.CallbackResponse{})
	})
	b.Handle(&listVote4, func(c *tb.Callback) {
		var msg = "Страница 4:\n\n"
		msg += proj4
		b.Edit(c.Message, msg, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: menuVote4})
		b.Respond(c, &tb.CallbackResponse{})
	})
	b.Handle(&vote1, func(c *tb.Callback) {
		var msg = "Вы за или против?\n\n"
		choseproj = proj1
		b.Edit(c.Message, msg, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: yesno})
		b.Respond(c, &tb.CallbackResponse{})
	})
	b.Handle(&vote2, func(c *tb.Callback) {
		var msg = "Вы за или против?\n\n"
		choseproj = proj2
		b.Edit(c.Message, msg, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: yesno})
		b.Respond(c, &tb.CallbackResponse{})
	})
	b.Handle(&vote3, func(c *tb.Callback) {
		var msg = "Вы за или против?\n\n"
		choseproj = proj3
		b.Edit(c.Message, msg, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: yesno})
		b.Respond(c, &tb.CallbackResponse{})
	})
	b.Handle(&vote4, func(c *tb.Callback) {
		var msg = "Вы за или против?\n\n"
		choseproj = proj4
		b.Edit(c.Message, msg, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: yesno})
		b.Respond(c, &tb.CallbackResponse{})
	})
	b.Handle(&yes, func(c *tb.Callback) {
		yesnores = "yes"
		var msg = "Вы проголосовали *за*, результаты можно будет посмотреть в личном кабинете"
		b.Edit(c.Message, msg, &tb.SendOptions{ParseMode: "Markdown"})
		b.Send(c.Sender, "Главное меню", &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{ResizeReplyKeyboard: true, ReplyKeyboard: mainMenu})
		b.Respond(c, &tb.CallbackResponse{})
	})
	b.Handle(&no, func(c *tb.Callback) {
		yesnores = "yes"
		var msg = "Вы проголосовали *против*, результаты можно будет посмотреть в личном кабинете"
		b.Edit(c.Message, msg, &tb.SendOptions{ParseMode: "Markdown"})
		b.Send(c.Sender, "Главное меню", &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{ResizeReplyKeyboard: true, ReplyKeyboard: mainMenu})
		b.Respond(c, &tb.CallbackResponse{})
	})
	b.Handle(&menu, func(c *tb.Callback) {
		yesnores = ""
		choseproj = ""
		b.Send(c.Sender, "Главное меню", &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{ResizeReplyKeyboard: true, ReplyKeyboard: mainMenu})
		b.Respond(c, &tb.CallbackResponse{})
	})
	b.Handle(&viewRes, func(c *tb.Callback) {
		var msg = "Список организаций, куда в проголосовали:\n\n"
		msg += "Организация: 1\n\n"
		msg += "Ваш голос: да\n\n"
		msg += "Тема голосования: 1\n\n"
		msg += "Завершен: да\n\n"
		msg += "Результат: 70 за и 30 против\n\n"
		b.Edit(c.Message, msg, &tb.SendOptions{ParseMode: "Markdown"})

		var msg2 = "Ваш публичный адрес: "
		// msg += addressPub
		msg2 += "\n\nВаш seed: "
		// msg += seed
		msg2 += "\n\nВаш баланс: "
		// msg += balance
		msg2 += " (РУБ)"
		msg2 += "\n\nТут вы можете посмотреть результаты прошлых голосований"

		b.Send(c.Sender, msg2, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: viewMenu})
		b.Respond(c, &tb.CallbackResponse{})
	})
	b.Start()
	// Обработчики на главное меню
}
