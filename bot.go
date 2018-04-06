package main

import (
	"log"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

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

	b.Handle("/hello", func(m *tb.Message) {
		b.Send(m.Sender, "hello world")
	})

	b.Start()
}
