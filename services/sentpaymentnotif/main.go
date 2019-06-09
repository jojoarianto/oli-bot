package sentpaymentnotif

import (
	"encoding/json"
	"github.com/jojoarianto/oli-bot/services/pushmessage/usecase"
	"net/http"
)

// DtoRequest data json from user input
type DtoRequest struct {
	UserID string `json:"user_id"`
}

// SendPaymentNotif for API handling for send message call from app
func SendPaymentNotif(w http.ResponseWriter, r *http.Request) {

	// == workflow ==
	// filter method only work on POST
	// get user_id & event_id
	// get payment where user_id & event_id
	// validate apakah sudah pernah dikirim notif
	// search on database table line_subscription where event_id
	// send push message to all row result
	// update status payment on database

	// filter method only work on POST
	switch r.Method {
	case http.MethodPost:
		var req DtoRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "400 - Bad Request", http.StatusBadRequest)
			return
		}

		if req.UserID == "" {
			http.Error(w, "400 - Bad Request", http.StatusBadRequest)
			return
		}
	default:
		http.Error(w, "405 - Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// get user_id & event_id

	// send push message
	usecase.PushMessage(
		"C0c5f795d51b9f447b966c79b772bd8ce",
		"https://cdn2.olimpiade.id/ero/payment-proof/payment-5ba066245b192836296c89f7.1541062118497434243015.jpg",
		"Irianto",
		"Rp. 100.000",
		"Bank Mandiri",
		"20-10-2019",
	)

	// update status payment on database
}
