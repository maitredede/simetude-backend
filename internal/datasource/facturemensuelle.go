package datasource

import "time"

type FactureMensuelle struct {
	KHWClient       float64 `json:"kwhClient"`
	AnneePrecedente float64 `json:"anneePrecedente"`
}

func (f FactureMensuelle) BesoinJournalier(m time.Month, year int) float64 {
	nDays := DaysInMonth(m, year)
	return f.KHWClient / float64(nDays)
}
