package repository

import (
	"github.com/jojoarianto/oli-bot/services/api/payment/model"
)

type GetPaymentRepository interface {
	GetPayment(userID string) (model.Payment, error)
}
