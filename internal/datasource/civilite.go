package datasource

type Civilite struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (c Civilite) getID() int {
	return c.ID
}

var (
	DataCivilite = build[Civilite](
		Civilite{ID: 1, Name: "Madame"},
		Civilite{ID: 2, Name: "Monsieur"},
		Civilite{ID: 3, Name: "Madame,Monsieur"},
		Civilite{ID: 4, Name: "Mademoiselle"},
		Civilite{ID: 5, Name: "SA"},
		Civilite{ID: 6, Name: "SARL"},
		Civilite{ID: 7, Name: "SAS"},
		Civilite{ID: 8, Name: "EURL"},
		Civilite{ID: 9, Name: "EARL"},
		Civilite{ID: 10, Name: "SCA"},
		Civilite{ID: 11, Name: "Société"},
	)
)
