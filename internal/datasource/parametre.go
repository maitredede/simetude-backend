package datasource

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

var (
	AllMonth []time.Month = []time.Month{
		time.January,
		time.February,
		time.March,
		time.April,
		time.May,
		time.June,
		time.July,
		time.August,
		time.September,
		time.October,
		time.November,
		time.December,
	}
)

func NewParametre() *Parametre {
	p := &Parametre{
		FactureClient: make(map[time.Month]FactureMensuelle),
	}
	p.RepartitionConsomation.p = p
	for _, m := range AllMonth {
		p.FactureClient[m] = FactureMensuelle{}
	}
	return p
}

type Parametre struct {
	// TypeClient     TypeClient `json:"typeClient"`
	// Civilite       Civilite   `json:"civilite"`
	// Nom            string     `json:"nom"`
	// Prenom         string     `json:"prenom"`
	// Adresse        string     `json:"adresse"`
	// CodePostal     string     `json:"codePostal"`
	// Ville          string     `json:"ville"`
	// TelMobile      string     `json:"telMobile"`
	// Email          string     `json:"email"`
	// Fournisseur    string     `json:"fournisseur"`
	// TypeAbonnement string     `json:"typeAbonnement"`
	// Piscine        bool       `json:"piscine"`
	// Climatiseur    bool       `json:"climatiseur"`

	FactureClient          map[time.Month]FactureMensuelle `json:"factureClient"`
	RepartitionConsomation RepartitionConsomation          `json:"repartitionConsomation"`
}

func (p Parametre) FactureAnnuelle() FactureMensuelle {
	total := FactureMensuelle{}
	for _, m := range AllMonth {
		if f, ok := p.FactureClient[m]; ok {
			total.KHWClient += f.KHWClient
			total.AnneePrecedente += f.AnneePrecedente
		}
	}
	return total
}

var _ json.Unmarshaler = (*Parametre)(nil)

func (p *Parametre) UnmarshalJSON(data []byte) error {
	type s Parametre
	var v s
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	*p = Parametre(v)
	p.RepartitionConsomation.p = p
	return nil
}

func (p *Parametre) Validate() error {
	var errs []error
	if p.RepartitionConsomation.p != p {
		errs = append(errs, errors.New("code: wrong data link for RepartitionConsomation"))
	}

	if len(p.FactureClient) != len(AllMonth) {
		errs = append(errs, fmt.Errorf("wrong length for FactureClient. len=%d expected=%d", len(p.FactureClient), len(AllMonth)))
	}
	for _, m := range AllMonth {
		_, ok := p.FactureClient[m]
		if !ok {
			errs = append(errs, fmt.Errorf("missing FactureClient for month %s (%d)", m, m))
		}
	}

	return errors.Join(errs...)
}
