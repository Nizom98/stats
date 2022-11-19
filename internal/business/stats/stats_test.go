package stats

import (
	"testing"

	"github.com/Nizom98/stats/internal/models"
	"github.com/stretchr/testify/assert"
)

//go:generate minimock -g -i github.com/Nizom98/stats/internal/models.StatsRepository -o ./repository_mock_test.go -n RepositoryMock

func TestNewEventCreated(t *testing.T) {
	repo := NewRepositoryMock(t)
	man := NewManager(repo)
	calledIncTotal, calledActivatedCalled := false, false

	repo.IncTotalMock.Set(func() {
		calledIncTotal = true
	})
	repo.ActivatedMock.Set(func() {
		calledActivatedCalled = true
	})
	repo.TransactionMock.Set(func(fn func(repo models.StatsRepository) error) (err error) {
		return fn(repo)
	})

	err := man.newEventCreated()

	assert.Nil(t, err)
	assert.True(t, calledIncTotal)
	assert.True(t, calledActivatedCalled)
}

func TestNewEventDeleted(t *testing.T) {
	repo := NewRepositoryMock(t)
	man := NewManager(repo)
	called := false

	repo.DeactivatedMock.Set(func() {
		called = true
	})
	repo.TransactionMock.Set(func(fn func(repo models.StatsRepository) error) (err error) {
		return fn(repo)
	})

	err := man.newEventDeleted()

	assert.Nil(t, err)
	assert.True(t, called)
}

func TestNewEventDeposited(t *testing.T) {
	repo := NewRepositoryMock(t)
	man := NewManager(repo)
	called := false
	expectAmount := float64(567)

	repo.DepositedMock.Set(func(amount float64) {
		assert.True(t, amount == expectAmount)
		called = true
	})
	repo.TransactionMock.Set(func(fn func(repo models.StatsRepository) error) (err error) {
		return fn(repo)
	})

	err := man.newEventDeposited(expectAmount)

	assert.Nil(t, err)
	assert.True(t, called)
}

func TestNewEventWithdrawn(t *testing.T) {
	repo := NewRepositoryMock(t)
	man := NewManager(repo)
	called := false
	expectAmount := float64(567)

	repo.WithdrawnMock.Set(func(amount float64) {
		assert.True(t, amount == expectAmount)
		called = true
	})
	repo.TransactionMock.Set(func(fn func(repo models.StatsRepository) error) (err error) {
		return fn(repo)
	})

	err := man.newEventWithdrawn(expectAmount)

	assert.Nil(t, err)
	assert.True(t, called)
}

func TestNewEventTransfered(t *testing.T) {
	repo := NewRepositoryMock(t)
	man := NewManager(repo)
	called := false
	expectAmount := float64(567)

	repo.TransferedMock.Set(func(amount float64) {
		assert.True(t, amount == expectAmount)
		called = true
	})
	repo.TransactionMock.Set(func(fn func(repo models.StatsRepository) error) (err error) {
		return fn(repo)
	})

	err := man.newEventTransfered(expectAmount)

	assert.Nil(t, err)
	assert.True(t, called)
}
