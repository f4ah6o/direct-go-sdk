package op

import "testing"

func TestParseSecretRef(t *testing.T) {
	vault, item, field, err := ParseSecretRef("op://path/to/direct_access_token")
	if err != nil {
		t.Fatal(err)
	}
	if vault != "Direct Bridge" || item != "account-a" || field != "direct_access_token" {
		t.Fatalf("unexpected parse: %q %q %q", vault, item, field)
	}
}

func TestParseSecretRefRejectsInvalid(t *testing.T) {
	if _, _, _, err := ParseSecretRef("not-a-ref"); err == nil {
		t.Fatalf("expected error")
	}
}
