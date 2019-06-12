package repository

import "github.com/jojoarianto/oli-bot/services/api/linesubscription/model"

type GetLineSubscription interface {
	GetByEventId(EventID string)([]model.LineSubscription, error)
}
