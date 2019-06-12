package sentpaymentnotif

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	linesubscription "github.com/jojoarianto/oli-bot/services/api/line/repository/mongodb"
	getpayment "github.com/jojoarianto/oli-bot/services/api/payment/repository/mongodb"
	"github.com/jojoarianto/oli-bot/services/pushmessage/config"
	"github.com/jojoarianto/oli-bot/services/pushmessage/usecase"
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

	var req DtoRequest

	// filter method only work on POST
	switch r.Method {
	case http.MethodPost:
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

	// get db connection
	db := config.Connect(os.Getenv("DB_CONNECTION"), os.Getenv("DB_NAME"))

	// get user_id & event_id
	repo := getpayment.NewGetPayment(db)
	payment, err := repo.GetPayment(req.UserID) // ada payment
	if err != nil {
		log.Print(err)
		return
	}

	// validate apakah sudah pernah dikirim notif atau sudah di approve
	if payment.PaymentStatus != 1 {
		// no need to send notification if status not equal with 1 (means only on waiting state)
		return
	}

	// search on database table line_subscription where event_id
	repoLineSubscription := linesubscription.NewGetLineSubscription(db)
	lineSubscritions, err := repoLineSubscription.GetByEventId("") // call get line_subscription
	if err != nil {
		log.Print(err)
		return
	}

	log.Print(lineSubscritions)

	// set up imageurl
	imageURL := fmt.Sprintf("https://cdn2.olimpiade.id/ero/payment-proof/%s", payment.PaymentProof)

	// send push message
	err = usecase.PushMessage(
		"C0c5f795d51b9f447b966c79b772bd8ce",
		imageURL,
		payment.AccountName,
		payment.Total,
		payment.BankName,
		payment.TransferAt,
	)
	if err != nil {
		return
	}

	// update status payment on database

}
