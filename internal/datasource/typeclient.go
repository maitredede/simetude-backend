package datasource

type TypeClient struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (c TypeClient) getID() int {
	return c.ID
}

var (
	DataTypeClient = build[TypeClient](
		TypeClient{ID: 1, Name: "Particulier"},
		TypeClient{ID: 2, Name: "Société"},
		TypeClient{ID: 3, Name: "Association"},
		TypeClient{ID: 4, Name: "Partenaire"},
	)
)
