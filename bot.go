package main

import (
	"log"
	"strconv"
	"time"

	"./mongo"
	"./userlogic"
	"./votes"
	"./waves"
	"github.com/vjeantet/jodaTime"

	mgo "gopkg.in/mgo.v2"
	tb "gopkg.in/tucnak/telebot.v2"
)

var session *mgo.Session

// var proj1 = "*Организация:* Фонд НеСкам\n\n*Адрес:* liwjbfiowbupiweubpwbep\n\n*Тема голосования:* Голосуем за то, чтобы провести новое ICO\n\n*Условия голосования:* иметь больше 100 noscum токенов\n\n*Время завершения голосования:* 9 Апреля, 00:01 (MSK)"
var proj1 []string
var proj2 = "*Организация:*  ------\n\n*Адрес:* ------\n\n*Тема голосования:* ------\n\n*Условия голосования:* ------\n\n*Время завершения голосования:* ------"
var proj3 = "*Организация:*  ---------\n\n*Адрес:* ------\n\n*Тема голосования:* ---------\n\n*Условия голосования:* ---------\n\n*Время завершения голосования:* ------"
var proj4 = "*Организация:*  ---------------\n\n*Адрес:* ---------\n\n*Тема голосования:* ------------\n\n*Условия голосования:* ------------\n\n*Время завершения голосования:* ------"
var choseproj = ""
var yesnores = ""
var enterName = false
var enterData = false
var golosTheme = ""
var golosData = ""

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
	session := mongo.ConnectToMongo()

	// Кнопки и меню
	mainData := tb.ReplyButton{Text: "💳 Мой кабинет"}
	votingData := tb.ReplyButton{Text: "📍 Текущие голосования"}
	createVote := tb.ReplyButton{Text: "🔖 Создать личное голосование"}
	mainMenu := [][]tb.ReplyButton{{mainData}, {votingData}, {createVote}}

	viewRes := tb.InlineButton{Unique: "viewres", Text: "Посмотреть, где я голосовал"}
	viewMy := tb.InlineButton{Unique: "viewMy", Text: "Созданные мной голосования"}
	viewMenu := [][]tb.InlineButton{{viewRes}, {viewMy}}

	// listVote1 := tb.InlineButton{Unique: "listvote1", Text: "Страница 1"}
	// listVote2 := tb.InlineButton{Unique: "listvote2", Text: "Страница 2"}
	// listVote3 := tb.InlineButton{Unique: "listvote3", Text: "Страница 3"}
	// listVote4 := tb.InlineButton{Unique: "listvote4", Text: "Страница 4"}

	// vote1 := tb.InlineButton{Unique: "vote1", Text: "Проголосовать за 1"}
	// vote2 := tb.InlineButton{Unique: "vote2", Text: "Проголосовать за 2"}
	// vote3 := tb.InlineButton{Unique: "vote3", Text: "Проголосовать за 3"}
	// vote4 := tb.InlineButton{Unique: "vote4", Text: "Проголосовать за 4"}

	menu := tb.InlineButton{Unique: "menu", Text: "Главное меню"}

	// menuVote1 := [][]tb.InlineButton{{vote1}, {listVote2, listVote3, listVote4}, {menu}}
	// menuVote2 := [][]tb.InlineButton{{vote2}, {listVote1, listVote3, listVote4}, {menu}}
	// menuVote3 := [][]tb.InlineButton{{vote3}, {listVote1, listVote2, listVote4}, {menu}}
	// menuVote4 := [][]tb.InlineButton{{vote4}, {listVote1, listVote2, listVote3}, {menu}}

	yes := tb.InlineButton{Unique: "yes", Text: "✅ За"}
	no := tb.InlineButton{Unique: "no", Text: "❌ Против"}
	yesno := [][]tb.InlineButton{
		{yes, no}, {menu}}

	cancel := tb.ReplyButton{Text: "Отмена"}
	cancelMenu := [][]tb.ReplyButton{{cancel}}

	voteyes := tb.ReplyButton{Text: "✅ Да, уверен"}
	voteno := tb.ReplyButton{Text: "❌ Нет"}
	voteyesno := [][]tb.ReplyButton{{voteyes, voteno}}

	voteyes2 := tb.ReplyButton{Text: "✅ Создать голосование"}
	voteno2 := tb.ReplyButton{Text: "❌ Нет, отменить"}
	voteyesno2 := [][]tb.ReplyButton{{voteyes2, voteno2}}
	// Кнопки и меню

	// Обработчики на главное меню
	b.Handle("/start", func(m *tb.Message) {
		var userID = strconv.Itoa(m.Sender.ID)

		if userlogic.Auth(session, userID) != true {
			var name = string(m.Sender.Username)
			var seed = waves.CreateSeed(userID, name)
			var address = waves.GetAddress(userID, seed)
			var msg = "Вы зарегистрированы в системе!\n\n"
			msg += "Ваш *Seed:* "
			msg += seed
			msg += "\n\n"
			msg += "Ваш *Address:* "
			msg += address
			b.Send(m.Sender, msg, &tb.SendOptions{DisableWebPagePreview: true, ParseMode: "Markdown"})
		}
		b.Send(m.Sender, "Главное меню", &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{ResizeReplyKeyboard: true, ReplyKeyboard: mainMenu})

	})
	b.Handle(&cancel, func(m *tb.Message) {
		golosTheme = ""
		golosData = ""
		enterData = false
		enterName = false
		b.Send(m.Sender, "Главное меню", &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{ResizeReplyKeyboard: true, ReplyKeyboard: mainMenu})
	})
	b.Handle(&mainData, func(m *tb.Message) {
		var userID = strconv.Itoa(m.Sender.ID)
		user := mongo.FindUser(session, userID)
		seed := waves.DecryptSeed(userID, user.EncryptedSeed)
		balance := waves.GetBalance(user.Address, "Waves")
		var msg = "Ваш публичный адрес: "
		msg += user.Address
		msg += "\n\nВаш seed: "
		msg += seed
		msg += "\n\nВаш баланс: "
		msg += balance
		// msg += " (РУБ)"
		msg += "\n\nТут вы можете посмотреть результаты прошлых голосований"
		b.Send(m.Sender, msg, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: viewMenu})
	})
	b.Handle(&viewRes, func(c *tb.Callback) {
		b.Edit(c.Message, &tb.SendOptions{ParseMode: "Markdown"}, "Тут вы можете посмотреть куда и как вы голосовали, а также узнать результат")
		b.Respond(c, &tb.CallbackResponse{})
	})
	b.Handle(&votingData, func(m *tb.Message) {
		var msg = "Список всех голосований:\n\n"

		votes := mongo.FindAllVotes(session)
		votesCount := len(votes)
		counter := 0
		if votesCount != 0 {
			for key := range votes {
				msg += "*Описание голосования:* \n" + votes[key].Description + "\n"
				msg += "*Чтобы проголосовать* нажмите на /"
				msg += "vote" + strconv.Itoa(counter) + "\n\n"
				counter++
			}

		} else {
			msg += "Голосования ещё не были созданы"
		}

		b.Send(m.Sender, msg, &tb.SendOptions{ParseMode: "Markdown"})
	})

	b.Handle("/vote0", func(m *tb.Message) {
		votes1 := mongo.FindAllVotes(session)
		startDate := jodaTime.Format("YYYY.MM.dd HH:mm", votes1[0].StartTime)
		endDate := jodaTime.Format("YYYY.MM.dd HH:mm", votes1[0].EndTime)
		choseproj = "0"
		msg := "*Описание голосования:* \n" + votes1[0].Description + "\n"
		msg += "*Начало:* \n" + (startDate) + "\n"
		msg += "*Окончание:* \n" + (endDate) + "\n"
		msg += "*Закончено ли:* \n"
		if votes1[0].End {
			msg += "ДА"
		} else {
			msg += "НЕТ"
		}
		b.Send(m.Sender, msg, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: yesno})
	})
	b.Handle("/vote1", func(m *tb.Message) {
		votes1 := mongo.FindAllVotes(session)
		startDate := jodaTime.Format("YYYY.MM.dd HH:mm", votes1[1].StartTime)
		endDate := jodaTime.Format("YYYY.MM.dd HH:mm", votes1[1].EndTime)
		choseproj = "1"
		msg := "*Описание голосования:* \n" + votes1[1].Description + "\n"
		msg += "*Начало:* \n" + (startDate) + "\n"
		msg += "*Окончание:* \n" + (endDate) + "\n"
		msg += "*Закончено ли:* \n"
		if votes1[1].End {
			msg += "ДА"
		} else {
			msg += "НЕТ"
		}
		b.Send(m.Sender, msg, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: yesno})
	})

	b.Handle("/vote2", func(m *tb.Message) {
		votes1 := mongo.FindAllVotes(session)
		startDate := jodaTime.Format("YYYY.MM.dd HH:mm", votes1[2].StartTime)
		endDate := jodaTime.Format("YYYY.MM.dd HH:mm", votes1[2].EndTime)
		choseproj = "2"
		msg := "*Описание голосования:* \n" + votes1[2].Description + "\n"
		msg += "*Начало:* \n" + (startDate) + "\n"
		msg += "*Окончание:* \n" + (endDate) + "\n"
		msg += "*Закончено ли:* \n"
		if votes1[2].End {
			msg += "ДА"
		} else {
			msg += "НЕТ"
		}
		b.Send(m.Sender, msg, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: yesno})
	})
	b.Handle("/vote3", func(m *tb.Message) {
		votes1 := mongo.FindAllVotes(session)
		startDate := jodaTime.Format("YYYY.MM.dd HH:mm", votes1[3].StartTime)
		endDate := jodaTime.Format("YYYY.MM.dd HH:mm", votes1[3].EndTime)
		choseproj = "3"
		msg := "*Описание голосования:* \n" + votes1[3].Description + "\n"
		msg += "*Начало:* \n" + (startDate) + "\n"
		msg += "*Окончание:* \n" + (endDate) + "\n"
		msg += "*Закончено ли:* \n"
		if votes1[3].End {
			msg += "ДА"
		} else {
			msg += "НЕТ"
		}
		b.Send(m.Sender, msg, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: yesno})
	})
	b.Handle("/vote4", func(m *tb.Message) {
		votes1 := mongo.FindAllVotes(session)
		startDate := jodaTime.Format("YYYY.MM.dd HH:mm", votes1[4].StartTime)
		endDate := jodaTime.Format("YYYY.MM.dd HH:mm", votes1[4].EndTime)
		choseproj = "4"
		msg := "*Описание голосования:* \n" + votes1[4].Description + "\n"
		msg += "*Начало:* \n" + (startDate) + "\n"
		msg += "*Окончание:* \n" + (endDate) + "\n"
		msg += "*Закончено ли:* \n"
		if votes1[4].End {
			msg += "ДА"
		} else {
			msg += "НЕТ"
		}
		b.Send(m.Sender, msg, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: yesno})
	})
	b.Handle("/vote5", func(m *tb.Message) {
		votes1 := mongo.FindAllVotes(session)
		startDate := jodaTime.Format("YYYY.MM.dd HH:mm", votes1[5].StartTime)
		endDate := jodaTime.Format("YYYY.MM.dd HH:mm", votes1[5].EndTime)
		choseproj = "5"
		msg := "*Описание голосования:* \n" + votes1[5].Description + "\n"
		msg += "*Начало:* \n" + (startDate) + "\n"
		msg += "*Окончание:* \n" + (endDate) + "\n"
		msg += "*Закончено ли:* \n"
		if votes1[5].End {
			msg += "ДА"
		} else {
			msg += "НЕТ"
		}
		b.Send(m.Sender, msg, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: yesno})
	})
	b.Handle("/vote6", func(m *tb.Message) {
		votes1 := mongo.FindAllVotes(session)
		startDate := jodaTime.Format("YYYY.MM.dd HH:mm", votes1[6].StartTime)
		endDate := jodaTime.Format("YYYY.MM.dd HH:mm", votes1[6].EndTime)
		choseproj = "6"
		msg := "*Описание голосования:* \n" + votes1[6].Description + "\n"
		msg += "*Начало:* \n" + (startDate) + "\n"
		msg += "*Окончание:* \n" + (endDate) + "\n"
		msg += "*Закончено ли:* \n"
		if votes1[6].End {
			msg += "ДА"
		} else {
			msg += "НЕТ"
		}
		b.Send(m.Sender, msg, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: yesno})
	})
	b.Handle("/vote7", func(m *tb.Message) {
		votes1 := mongo.FindAllVotes(session)
		startDate := jodaTime.Format("YYYY.MM.dd HH:mm", votes1[7].StartTime)
		endDate := jodaTime.Format("YYYY.MM.dd HH:mm", votes1[7].EndTime)
		choseproj = "7"
		msg := "*Описание голосования:* \n" + votes1[7].Description + "\n"
		msg += "*Начало:* \n" + (startDate) + "\n"
		msg += "*Окончание:* \n" + (endDate) + "\n"
		msg += "*Закончено ли:* \n"
		if votes1[7].End {
			msg += "ДА"
		} else {
			msg += "НЕТ"
		}
		b.Send(m.Sender, msg, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: yesno})
	})
	b.Handle("/vote8", func(m *tb.Message) {
		votes1 := mongo.FindAllVotes(session)
		startDate := jodaTime.Format("YYYY.MM.dd HH:mm", votes1[8].StartTime)
		endDate := jodaTime.Format("YYYY.MM.dd HH:mm", votes1[8].EndTime)
		choseproj = "8"
		msg := "*Описание голосования:* \n" + votes1[8].Description + "\n"
		msg += "*Начало:* \n" + (startDate) + "\n"
		msg += "*Окончание:* \n" + (endDate) + "\n"
		msg += "*Закончено ли:* \n"
		if votes1[8].End {
			msg += "ДА"
		} else {
			msg += "НЕТ"
		}
		b.Send(m.Sender, msg, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: yesno})
	})
	b.Handle("/vote9", func(m *tb.Message) {
		votes1 := mongo.FindAllVotes(session)
		startDate := jodaTime.Format("YYYY.MM.dd HH:mm", votes1[9].StartTime)
		endDate := jodaTime.Format("YYYY.MM.dd HH:mm", votes1[9].EndTime)
		choseproj = "9"
		msg := "*Описание голосования:* \n" + votes1[9].Description + "\n"
		msg += "*Начало:* \n" + (startDate) + "\n"
		msg += "*Окончание:* \n" + (endDate) + "\n"
		msg += "*Закончено ли:* \n"
		if votes1[9].End {
			msg += "ДА"
		} else {
			msg += "НЕТ"
		}
		b.Send(m.Sender, msg, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: yesno})
	})
	b.Handle("/vote10", func(m *tb.Message) {
		votes1 := mongo.FindAllVotes(session)
		startDate := jodaTime.Format("YYYY.MM.dd HH:mm", votes1[10].StartTime)
		endDate := jodaTime.Format("YYYY.MM.dd HH:mm", votes1[10].EndTime)
		choseproj = "10"
		msg := "*Описание голосования:* \n" + votes1[10].Description + "\n"
		msg += "*Начало:* \n" + (startDate) + "\n"
		msg += "*Окончание:* \n" + (endDate) + "\n"
		msg += "*Закончено ли:* \n"
		if votes1[10].End {
			msg += "ДА"
		} else {
			msg += "НЕТ"
		}
		b.Send(m.Sender, msg, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{InlineKeyboard: yesno})
	})

	b.Handle(&yes, func(c *tb.Callback) {
		yesnores = "yes"
		userID := strconv.Itoa(c.Sender.ID)
		user := mongo.FindUser(session, userID)
		voters := mongo.FindAllVoters(session)
		var msg string
		check := 0
		num, _ := strconv.Atoi(choseproj)
		for key := range voters {
			if voters[key].Address == user.Address && voters[key].Num == num {
				check++
			}
		}

		if check != 1 {
			votes.Vote(choseproj, user.Address, "1")
			msg += "Вы проголосовали *за*, результаты можно будет посмотреть в личном кабинете"
		} else {
			msg += "Вы уже голосовали"
		}

		b.Edit(c.Message, msg, &tb.SendOptions{ParseMode: "Markdown"})
		b.Send(c.Sender, "Главное меню", &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{ResizeReplyKeyboard: true, ReplyKeyboard: mainMenu})
		b.Respond(c, &tb.CallbackResponse{})
	})
	b.Handle(&no, func(c *tb.Callback) {
		yesnores = "yes"
		userID := strconv.Itoa(c.Sender.ID)
		user := mongo.FindUser(session, userID)
		voters := mongo.FindAllVoters(session)
		var msg string
		check := 0
		num, _ := strconv.Atoi(choseproj)
		for key := range voters {
			if voters[key].Address == user.Address && voters[key].Num == num {
				check++
			}
		}

		if check != 1 {
			votes.Vote(choseproj, user.Address, "0")
			msg += "Вы проголосовали *против*, результаты можно будет посмотреть в личном кабинете"
		} else {
			msg += "Вы уже голосовали"
		}

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
	b.Handle(&createVote, func(m *tb.Message) {
		var msg = "На текущий момент поддерживаются лишь голосования типа: да / нет:\n\n"
		msg += "На какую тематику будет ваше голосование? (Опишите о чем будет голосование, чтобы участники могли ответить либо да, либо нет"
		enterName = true
		b.Send(m.Sender, msg, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{ResizeReplyKeyboard: true, ReplyKeyboard: cancelMenu})
		b.Handle(tb.OnText, func(m *tb.Message) {
			if enterName {
				golosTheme = m.Text
				var msg1 = "Вы уверены, что хотите создать голосование с данной тематикой?"
				b.Send(m.Sender, msg1, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{ResizeReplyKeyboard: true, ReplyKeyboard: voteyesno})
				enterName = false
			}
		})
	})
	b.Handle(&voteyes, func(m *tb.Message) {
		var msg = "Ваша тематика "
		msg += golosTheme
		enterData = true
		msg += "\n\nТеперь введите длительность голосования в *часах* *(Не больше 10 часов)*"
		b.Send(m.Sender, msg, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{ResizeReplyKeyboard: true, ReplyKeyboard: cancelMenu})
		b.Handle(tb.OnText, func(m *tb.Message) {
			if enterData {
				var msg1 = "Вы уверены, что хотите создать голосование на "
				msg1 += m.Text
				golosData = m.Text
				msg1 += " ч."
				b.Send(m.Sender, msg1, &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{ResizeReplyKeyboard: true, ReplyKeyboard: voteyesno2})
				enterData = false
			}
		})
	})
	b.Handle(&voteyes2, func(m *tb.Message) {
		userID := strconv.Itoa(m.Sender.ID)
		votes.CreateVote(userID, golosTheme, golosData)
		b.Send(m.Sender, "Вы успешно создали голосование, его можно посмотреть в *текущих голосованиях* или в *моих голосованиях* в *моем кабинете*", &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{ResizeReplyKeyboard: true, ReplyKeyboard: mainMenu})
	})
	b.Handle(&voteno2, func(m *tb.Message) {
		golosData = ""
		enterData = false
		enterName = false
		b.Send(m.Sender, "Главное меню", &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{ResizeReplyKeyboard: true, ReplyKeyboard: mainMenu})
	})
	b.Handle(&voteno, func(m *tb.Message) {
		golosTheme = ""
		golosData = ""
		enterData = false
		enterName = false
		b.Send(m.Sender, "Главное меню", &tb.SendOptions{ParseMode: "Markdown"}, &tb.ReplyMarkup{ResizeReplyKeyboard: true, ReplyKeyboard: mainMenu})
	})
	b.Handle(&viewMy, func(c *tb.Callback) {
		var msg = "Список созданных вами голосований: \n\n"
		msg += "Голосование 1\n"
		msg += "Тематика: \n"
		msg += "Сколько времени до конца: \n"
		msg += "Завершен: нет\n"
		msg += "Результаты: "
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
