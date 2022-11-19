package repository

type stats struct {
	total      int64
	active     int64
	inactive   int64
	deposited  float64
	withdrawn  float64
	transfered float64
}

func (s *stats) Total() int64 {
	return s.total
}

func (s *stats) Active() int64 {
	return s.active
}

func (s *stats) Inactive() int64 {
	return s.inactive
}

func (s *stats) Deposited() float64 {
	return s.deposited
}

func (s *stats) Withdrawn() float64 {
	return s.withdrawn
}

func (s *stats) Transfered() float64 {
	return s.transfered
}
