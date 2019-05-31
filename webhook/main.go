package webhook

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/line/line-bot-sdk-go/linebot"
)

// Webhook is method callback for line webhook
func Webhook(w http.ResponseWriter, r *http.Request) {

	bot, err := linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}

	events, err := bot.ParseRequest(r)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:

				// accept only command when first letter is /
				if message.Text[0] == '/' {

					inputArr := strings.Fields(message.Text)

					switch inputArr[0] { // filter for first word
					case "/register":
						str := "Hello you are in registering"
						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(str)).Do(); err != nil {
							log.Print(err)
						}
					case "/profile":
						userID := event.Source.UserID
						if userID != "" {
							str := fmt.Sprintf("Your id is %s", userID)
							if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(str)).Do(); err != nil {
								log.Print(err)
							}
						}
					case "/group":
						groupID := event.Source.GroupID
						if groupID != "" {
							str := fmt.Sprintf("Your group id is %s", groupID)
							if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(str)).Do(); err != nil {
								log.Print(err)
							}
						}
					case "/room":
						roomID := event.Source.RoomID
						if roomID != "" {
							str := fmt.Sprintf("Your room id is %s", roomID)
							if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(str)).Do(); err != nil {
								log.Print(err)
							}
						}
					default:
						str := "Perintah tidak ditemukan, silakan lihat daftar perintah dengan mengetikkan /help"
						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(str)).Do(); err != nil {
							log.Print(err)
						}
					}

				} else {
					// DO NOTHING IS GOOD
				}

			}
		}
	}
}
