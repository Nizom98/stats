package rest

import (
	"encoding/json"
	"github.com/Nizom98/stats/internal/models"
	"net/http"
)

func NewHandler(manStats models.StatsManager) (*Handler, error) {
	return &Handler{
		manStats: manStats,
	}, nil
}

func (h *Handler) StatsHandler(w http.ResponseWriter, req *http.Request) {
	stats := h.manStats.Stats()
	resp := &StatsData{
		Total:      stats.Total(),
		Active:     stats.Active(),
		Inactive:   stats.Inactive(),
		Deposited:  stats.Deposited(),
		Withdrawn:  stats.Withdrawn(),
		Transfered: stats.Transfered(),
	}

	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(
		&StatsResponse{
			Wallet: resp,
		},
	)
	panic(err)
}

func printError(w http.ResponseWriter, err string, status int) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(
		&StatusResponse{
			Success:    false,
			ErrMessage: err,
		},
	)
}
