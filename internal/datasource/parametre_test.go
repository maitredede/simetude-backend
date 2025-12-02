package datasource

import (
	_ "embed"
	"encoding/json"
	"testing"
)

var (
	//go:embed parametre_test.json
	testData []byte
)

func TestUnmarshal(t *testing.T) {
	var p Parametre
	if err := json.Unmarshal(testData, &p); err != nil {
		t.Fatal(err)
	}
	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
