package datasource

type RepartitionConsomation struct {
	p *Parametre

	TravailExterieur  bool    `json:"travailExterieur"`
	ConsoJourPourcent float64 `json:"consoJourPourcent"`
	ConsoNuitPourcent float64 `json:"consoNuitPourcent"`
	SolaireExistant   bool    `json:"solaireExistant"`
}

func (r RepartitionConsomation) KWHAnnuelJour() float64 {
	return r.ConsoJourPourcent * r.p.FactureAnnuelle().KHWClient
}

func (r RepartitionConsomation) KWHAnnuelNuit() float64 {
	return r.ConsoNuitPourcent * r.p.FactureAnnuelle().KHWClient
}
