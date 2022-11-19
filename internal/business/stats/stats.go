package stats

import (
	"fmt"

	"github.com/Nizom98/stats/internal/models"
	log "github.com/sirupsen/logrus"
)

const (
	eventWalletCreated    = "Wallet_Created"
	eventWalletDeleted    = "Wallet_Deleted"
	eventWalletDeposited  = "Wallet_Deposited"
	eventWalletWithdrawn  = "Wallet_Withdrawn"
	eventWalletTransfered = "Wallet_Transfered"
)

type manager struct {
	repo models.StatsRepository
}

// Stats получаем данные о статистике
func (man *manager) Stats() models.Stats {
	return man.repo.Stats()
}

// NewManager конструктор менеджера кошельков
func NewManager(repo models.StatsRepository) *manager {
	return &manager{
		repo: repo,
	}
}

// newEventCreated обрабатываем событие о создании кошелька
func (man *manager) newEventCreated() error {
	errTx := man.repo.Transaction(func(repo models.StatsRepository) error {
		repo.IncTotal()
		repo.Activated()
		return nil
	})
	return errTx
}

// newEventDeposited обрабатываем событие о деактивации кошелька
func (man *manager) newEventDeleted() error {
	errTx := man.repo.Transaction(func(repo models.StatsRepository) error {
		repo.Deactivated()
		return nil
	})
	return errTx
}

// newEventDeposited обрабатываем событие о пополнении
func (man *manager) newEventDeposited(amount float64) error {
	errTx := man.repo.Transaction(func(repo models.StatsRepository) error {
		repo.Deposited(amount)
		return nil
	})
	return errTx
}

// newEventWithdrawn обрабатываем событие о снятии
func (man *manager) newEventWithdrawn(amount float64) error {
	errTx := man.repo.Transaction(func(repo models.StatsRepository) error {
		repo.Withdrawn(amount)
		return nil
	})
	return errTx
}

// newEventTransfered обрабатываем событие о переводе
func (man *manager) newEventTransfered(amount float64) error {
	errTx := man.repo.Transaction(func(repo models.StatsRepository) error {
		repo.Transfered(amount)
		return nil
	})
	return errTx
}

// EventHandler обрабатывает события
func (man *manager) EventHandler(event *models.Event) {
	if event == nil {
		log.Errorf("unexpected empty event")
		return
	}

	log.Infof("start consume event: type %s, amount %f", event.Type, event.Amount)
	var err error
	switch event.Type {
	case eventWalletCreated:
		err = man.newEventCreated()
	case eventWalletDeleted:
		err = man.newEventDeleted()
	case eventWalletDeposited:
		err = man.newEventDeposited(event.Amount)
	case eventWalletWithdrawn:
		err = man.newEventWithdrawn(event.Amount)
	case eventWalletTransfered:
		err = man.newEventTransfered(event.Amount)
	default:
		err = fmt.Errorf("unknown event type: %s", event.Type)
	}

	if err != nil {
		log.Errorf("err while consume msg: %s", err.Error())
		return
	}

	log.Infof("end consume event: type %s", event.Type)
}
