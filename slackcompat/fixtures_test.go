package slackcompat

import (
	"encoding/json"
	"os"
	"testing"
)

type fixtureCase struct {
	Name    string                 `json:"name"`
	Method  string                 `json:"method,omitempty"`
	Request map[string]interface{} `json:"request,omitempty"`
	Expect  map[string]interface{} `json:"expect,omitempty"`
	Event   map[string]interface{} `json:"event,omitempty"`
}

func TestFixtureOracleLoads(t *testing.T) {
	data, err := os.ReadFile("testdata/web_api_fixtures.json")
	if err != nil {
		t.Fatal(err)
	}
	var fixtures []fixtureCase
	if err := json.Unmarshal(data, &fixtures); err != nil {
		t.Fatal(err)
	}
	if len(fixtures) < 3 {
		t.Fatalf("fixtures len = %d, want at least 3", len(fixtures))
	}
	for _, fixture := range fixtures {
		if fixture.Name == "" {
			t.Fatalf("fixture without name: %+v", fixture)
		}
		if fixture.Method == "" && fixture.Event == nil {
			t.Fatalf("fixture %q needs method or event", fixture.Name)
		}
	}
}
