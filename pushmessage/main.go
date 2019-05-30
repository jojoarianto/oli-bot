package pushmessage

import (
	"log"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
)

// PushMessage method to send bc
func PushMessage(w http.ResponseWriter, r *http.Request) {
	bot, err := linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// read input from request
	
	jsonString := PaymentNotifTemplate(
		"https://cdn2.olimpiade.id/ero/payment-proof/payment-5ba066245b192836296c89f7.1541062118497434243015.jpg",
		"Irianto",
		"Rp. 100.000",
		"Bank Mandiri",
		"20-10-2019",
	)
	contents, err := linebot.UnmarshalFlexMessageJSON([]byte(jsonString))
	if err != nil {
		log.Print(err)
	}

	// looping each subscription
	if _, err := bot.PushMessage("Ue2e068023850bff76bf462e103509fad", linebot.NewFlexMessage("New payment need approval", contents)).Do(); err != nil {
		log.Print(err)
	}
}
