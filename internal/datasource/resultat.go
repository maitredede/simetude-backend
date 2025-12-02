package datasource

import "time"

func Calcule(p Parametre) Resultat {
	year := time.Now().Year()

	r := Resultat{}

	r.TarifKWh = 39
	r.AbonnementKvaMois = 555
	r.RedevanceComptage = 641
	r.TauxRedevanceCommunale = 9
	r.TauxTGC = 3

	r.ConsoAnneePrecedente = p.FactureAnnuelle().AnneePrecedente
	r.ConsoAnneeCourante = p.FactureAnnuelle().KHWClient
	r.MoyenneAnnee = (r.ConsoAnneePrecedente + r.ConsoAnneeCourante) / 2 // moyenne
	r.BesoinsJournaliers = 0
	for _, m := range AllMonth {
		r.BesoinsJournaliers += p.FactureClient[m].BesoinJournalier(m, year)
	}
	r.BesoinsJournaliers = r.BesoinsJournaliers / float64(len(AllMonth))

	return r
}

type Resultat struct {
	// InfosClient            any
	// TypeAbonnement         string
	// PuissanceSouscrite     string
	TarifKWh               float64
	AbonnementKvaMois      float64
	RedevanceComptage      float64
	TauxRedevanceCommunale float64
	TauxTGC                float64

	ConsoAnneePrecedente float64
	ConsoAnneeCourante   float64
	MoyenneAnnee         float64

	BesoinsJournaliers float64
}
