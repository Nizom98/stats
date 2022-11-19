package repository

import (
	"example.com/m/v2/internal/models"
	"sync"
)

// StatsRepository ...
type StatsRepository struct {
	muStats *sync.RWMutex
	stats   *stats
}

// NewRepo конструктор репозитория
func NewRepo() *StatsRepository {
	return &StatsRepository{
		muStats: new(sync.RWMutex),
		stats:   &stats{},
	}
}

// Transaction для конкурентной записи в хранилище.
func (repo *StatsRepository) Transaction(fn func(repo *StatsRepository) error) error {
	repo.muStats.Lock()
	defer repo.muStats.Unlock()

	return fn(repo)
}

// IncTotal увеличиваем общее количество кошельков на 1
func (repo *StatsRepository) IncTotal() {
	repo.stats.total++
}

// Activated увеличиваем общее количество активных кошельков на 1
func (repo *StatsRepository) Activated() {
	repo.stats.active++
}

// Deactivated увеличиваем общее количество неактивных кошельков на 1
func (repo *StatsRepository) Deactivated() {
	repo.stats.inactive++
}

// Deposited складываем к общей сумме пополнений amount
func (repo *StatsRepository) Deposited(amount float64) {
	repo.stats.deposited += amount
}

// Withdrawn складываем к общей сумме снятий amount
func (repo *StatsRepository) Withdrawn(amount float64) {
	repo.stats.withdrawn += amount
}

// Transfered складываем к общей сумме переводов amount
func (repo *StatsRepository) Transfered(amount float64) {
	repo.stats.transfered += amount
}

// Stats получаем данные о статистике
func (repo *StatsRepository) Stats() models.Stats {
	return repo.stats
}
