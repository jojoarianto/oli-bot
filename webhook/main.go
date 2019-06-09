package webhook

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/jojoarianto/oli-bot/webhook/route"
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

						// call the register prosses handler
						RegisterHandler(event, bot, inputArr)

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

// RegisterHandler for handling registration line (group, individe or room) to subscription
// inputArr is equal with {"email", "password", "line_id", "type"}
func RegisterHandler(event *linebot.Event, bot *linebot.Client, inputArr []string) {

	const (
		Group    = "group"
		Room     = "room"
		Personal = "personal"
	)

	var (
		lineAccountType = Personal // set default to personal
		lineAccountID   = ""       // set default empty id
	)

	//	cek apakah dia register di dalam group
	groupID := event.Source.GroupID
	if groupID != "" {
		lineAccountType = Group
		lineAccountID = groupID
	}

	// cek apakah dia register di dalam room
	roomID := event.Source.RoomID
	if roomID != "" {
		lineAccountType = Room
		lineAccountID = roomID
	}

	// jika tidak semua apakah dia menambahkan di dalam private message mode
	userID := event.Source.UserID
	if groupID == "" && roomID == "" {
		lineAccountType = Personal
		lineAccountID = userID
	}

	// inject for fill lineID & LineAccountType
	inputArr = append(inputArr, lineAccountID, lineAccountType)

	// sent regiter with type status
	str := route.Register(inputArr)
	if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(str)).Do(); err != nil {
		log.Print(err)
	}
}
