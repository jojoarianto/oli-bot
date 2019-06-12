package mongodb

import (
	"github.com/jojoarianto/oli-bot/services/api/linesubscription/model"
	"github.com/jojoarianto/oli-bot/services/api/linesubscription/repository"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type mongodbGetLineSubscriptionRepository struct {
	Conn *mgo.Database
}

// NewGetLineSubscription contructor repo
func NewGetLineSubscription(Conn *mgo.Database) repository.GetLineSubscription {
	return &mongodbGetLineSubscriptionRepository{Conn}
}

// GetByEventId method to get linesubscription by event id
func (m *mongodbGetLineSubscriptionRepository) GetByEventId(EventID string) ([]model.LineSubscription, error) {
	var result []model.LineSubscription
	err := m.Conn.C("users").Find(bson.M{}).All(&result)
	return result, err
}
