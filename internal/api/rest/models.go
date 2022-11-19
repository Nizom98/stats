package rest

import "example.com/m/v2/internal/models"

type Handler struct {
	manStats models.StatsManager
}

type StatsData struct {
	Total      int64   `json:"total"`
	Active     int64   `json:"active"`
	Inactive   int64   `json:"inactive"`
	Deposited  float64 `json:"deposited"`
	Withdrawn  float64 `json:"withdrawn"`
	Transfered float64 `json:"transfered"`
}

type StatsResponse struct {
	Wallet *StatsData `json:"wallet"`
}

type StatusResponse struct {
	Success    bool        `json:"success"`
	ErrMessage string      `json:"err_message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}
