package model

type LineSubscription struct {
	ID      string `json:"id"`
	EventID string `json:"event_id"`
	LineID  string `json:"line_id"`
	Type    string `json:"type"`
	UserID  string `json:"user_id"`
}
