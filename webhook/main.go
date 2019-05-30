package webhook

import (
	"log"
	"net/http"
	"os"

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

				switch message.Text {
				case "/profile":
					userID := event.Source.UserID
					if userID != "" {
						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(userID)).Do(); err != nil {
							log.Print(err)
						}
					}
				case "/group":
					groupID := event.Source.GroupID
					if groupID != "" {
						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(groupID)).Do(); err != nil {
							log.Print(err)
						}
					}
				case "/room":
					roomID := event.Source.RoomID
					if roomID != "" {
						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(roomID)).Do(); err != nil {
							log.Print(err)
						}
					}
				default:
					// not implement yet
				}

			}
		}
	}
}
