package repository

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIncTotal(t *testing.T) {
	repo := fakeStatsRepo()
	initValue := int64(57)
	expectValue := initValue + 1
	repo.stats.total = initValue

	repo.IncTotal()
	assert.True(t, repo.stats.total == expectValue)
}

func TestActivated(t *testing.T) {
	repo := fakeStatsRepo()
	initValue := int64(57)
	expectValue := initValue + 1
	repo.stats.active = initValue

	repo.Activated()
	assert.True(t, repo.stats.active == expectValue)
}

func TestDeactivated(t *testing.T) {
	repo := fakeStatsRepo()
	initValue := int64(57)
	expectValue := initValue + 1
	repo.stats.inactive = initValue

	repo.Deactivated()
	assert.True(t, repo.stats.inactive == expectValue)
}

func TestDeposited(t *testing.T) {
	repo := fakeStatsRepo()
	initValue, amount := float64(57), float64(765)
	repo.stats.deposited = initValue

	repo.Deposited(amount)
	assert.True(t, repo.stats.deposited == initValue+amount)
}

func TestWithdrawn(t *testing.T) {
	repo := fakeStatsRepo()
	initValue, amount := float64(43), float64(98)
	repo.stats.withdrawn = initValue

	repo.Withdrawn(amount)
	assert.True(t, repo.stats.withdrawn == initValue+amount)
}

func TestTransfered(t *testing.T) {
	repo := fakeStatsRepo()
	initValue, amount := float64(43), float64(98)
	repo.stats.transfered = initValue

	repo.Transfered(amount)
	assert.True(t, repo.stats.transfered == initValue+amount)
}

func fakeStatsRepo() *StatsRepository {
	return &StatsRepository{
		stats: &stats{},
	}
}
