package models

type Event struct {
	Type   string  `json:"type"`
	Amount float64 `json:"amount"`
}
