package repository

import "github.com/jojoarianto/oli-bot/services/api/line/model"

type GetLineSubscription interface {
	GetByEventId(EventID string) ([]model.LineSubscription, error)
}
