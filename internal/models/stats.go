package models

type Stats interface {
	Total() int64
	Active() int64
	Inactive() int64
	Deposited() float64
	Withdrawn() float64
	Transfered() float64
}

type StatsManager interface {
	EventHandler(event *Event)
	Stats() Stats
}

type StatsRepository interface {
	Transaction(fn func(repo StatsRepository) error) error
	IncTotal()
	Activated()
	Deactivated()
	Deposited(amount float64)
	Withdrawn(amount float64)
	Transfered(amount float64)
	Stats() Stats
}
