package mongodb

import (
	"gopkg.in/mgo.v2"
	"log"
	"testing"
)

// Connect make connection
func Connect(strUri, dbName string) *mgo.Database {
	session, err := mgo.Dial(strUri)
	if err != nil {
		log.Fatal(err)
	}
	return session.DB(dbName)
}

func Test_mongodbPaymentRepository_GetPayment(t *testing.T) {
	conn := Connect("localhost~", "sibiti")

	repo := NewGetPayment(conn)
	payment, err := repo.GetPayment("5cf32d455b192833a81e153d") // ada payment
	//payment, err := repo.GetPayment("5cfbdc605b19283738016a58") // tidak ada payment

	if err != nil {
		t.Error(err)
	}

	t.Log(payment)
	t.Log(payment.PaymentStatus)
}
