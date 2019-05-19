package main

import (
	"fmt"
	"log"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {
	bot, err := linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}

	jsonString := PaymentNotifTemplate("https://cdn2.olimpiade.id/ero/payment-proof/payment-5ba066245b192836296c89f7.1541062118497434243015.jpg", "Irianto")
	contents, err := linebot.UnmarshalFlexMessageJSON([]byte(jsonString))
	if err != nil {
		log.Print(err)
	}

	if _, err := bot.PushMessage("Ue2e068023850bff76bf462e103509fad", linebot.NewFlexMessage("New payment need approval", contents)).Do(); err != nil {
		log.Print(err)
	}
}

// PaymentNotifTemplate tempate for notif payment message builder
func PaymentNotifTemplate(imageURL string, name string) string {
	template := `
	{
		"type": "bubble",
		"hero": {
		  "type": "image",
		  "url": "%s",
		  "size": "full",
		  "aspectRatio": "20:13",
		  "aspectMode": "cover",
		  "action": {
			"type": "uri",
			"uri": "%s"
		  }
		},
		"body": {
		  "type": "box",
		  "layout": "vertical",
		  "contents": [
			{
			  "type": "text",
			  "text": "%s",
			  "weight": "bold",
			  "size": "xl"
			},
			{
			  "type": "box",
			  "layout": "vertical",
			  "margin": "lg",
			  "spacing": "sm",
			  "contents": [
				{
				  "type": "box",
				  "layout": "baseline",
				  "spacing": "sm",
				  "contents": [
					{
					  "type": "text",
					  "text": "Amount",
					  "color": "#aaaaaa",
					  "size": "sm",
					  "flex": 1
					},
					{
					  "type": "text",
					  "text": "Rp. 150.000,-",
					  "wrap": true,
					  "color": "#666666",
					  "size": "sm",
					  "flex": 5
					}
				  ]
				},
				{
				  "type": "box",
				  "layout": "baseline",
				  "spacing": "sm",
				  "contents": [
					{
					  "type": "text",
					  "text": "Bank",
					  "color": "#aaaaaa",
					  "size": "sm",
					  "flex": 1
					},
					{
					  "type": "text",
					  "text": "Bank Mandiri",
					  "wrap": true,
					  "color": "#666666",
					  "size": "sm",
					  "flex": 5
					}
				  ]
				},
				{
				  "type": "box",
				  "layout": "baseline",
				  "spacing": "sm",
				  "contents": [
					{
					  "type": "text",
					  "text": "Date",
					  "color": "#aaaaaa",
					  "size": "sm",
					  "flex": 1
					},
					{
					  "type": "text",
					  "text": "29-09-2018",
					  "wrap": true,
					  "color": "#666666",
					  "size": "sm",
					  "flex": 5
					}
				  ]
				}
			  ]
			}
		  ]
		},
		"footer": {
		  "type": "box",
		  "layout": "vertical",
		  "spacing": "sm",
		  "contents": [
			{
			  "type": "button",
			  "style": "link",
			  "height": "sm",
			  "action": {
				"type": "uri",
				"label": "PERGI KE PORTAL",
				"uri": "https://portal.olimpiade.id/kelola-payment/waiting"
			  }
			},
			{
			  "type": "spacer",
			  "size": "sm"
			}
		  ],
		  "flex": 0
		}
	}`

	jsonString := fmt.Sprintf(template, imageURL, imageURL, name)
	return jsonString
}
