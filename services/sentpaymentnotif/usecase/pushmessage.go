package usecase

import (
	"github.com/jojoarianto/oli-bot/services/pushmessage/template"
	"log"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
)

// PushMessage method to send bc
func PushMessage(to, imageUrl, name, nominal, bank, strDate string) {

	bot, err := linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// read input from request
	jsonString := template.PaymentNotifTemplate(imageUrl, name, nominal, bank, strDate)

	contents, err := linebot.UnmarshalFlexMessageJSON([]byte(jsonString))
	if err != nil {
		log.Print(err)
	}

	// looping each subscription
	if _, err := bot.PushMessage(to, linebot.NewFlexMessage("New payment need approval", contents)).Do(); err != nil {
		log.Print(err)
	}
}
