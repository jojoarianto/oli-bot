package mongodb

import (
	"github.com/jojoarianto/oli-bot/services/api/payment/model"
	"github.com/jojoarianto/oli-bot/services/api/payment/repository"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type mongodbPaymentRepository struct {
	Conn *mgo.Database
}

func NewGetPayment(Conn *mgo.Database) repository.GetPaymentRepository {
	return &mongodbPaymentRepository{Conn}
}

func (m *mongodbPaymentRepository) GetPayment(userID string) (model.Payment, error) {
	var (
		usr     model.User
		payment model.Payment
	)
	err := m.Conn.C("users").FindId(bson.ObjectIdHex(userID)).One(&usr)
	if err != nil {
		return payment, err
	}

	// midleware to check apakah payment data di temukan
	if usr.Payment["payment_status"] == nil {
		return payment, nil
	}

	// handling for null value
	// coz some of payment value may be doesn't exist
	if usr.Payment["payment_proof"] != nil {
		payment.PaymentProof = usr.Payment["payment_proof"].(string)
	}
	if usr.Payment["payment_proof_thumbnail"] != nil {
		payment.PaymentProofThumbnail = usr.Payment["payment_proof_thumbnail"].(string)
	}
	if usr.Payment["account_name"] != nil {
		payment.AccountName = usr.Payment["account_name"].(string)
	}
	if usr.Payment["bank_name"] != nil {
		payment.BankName = usr.Payment["bank_name"].(string)
	}
	if usr.Payment["transfered_at"] != nil {
		payment.TransferAt= usr.Payment["transfered_at"].(string)
	}
	if usr.Payment["total"] != nil {
		payment.Total = usr.Payment["total"].(string)
	}
	// for int value (only on payment status)
	if usr.Payment["payment_status"] != nil || usr.Payment["payment_status"] != 0 {
		payment.PaymentStatus = usr.Payment["payment_status"].(int)
	}

	return payment, err
}
