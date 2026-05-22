package op

import (
	"strings"
	"testing"
)

func TestParseSecretRef(t *testing.T) {
	vault, item, field, err := ParseSecretRef("op://path/to/direct_access_token")
	if err != nil {
		t.Fatal(err)
	}
	if vault != "path" || item != "to" || field != "direct_access_token" {
		t.Fatalf("unexpected parse: %q %q %q", vault, item, field)
	}
}

func TestParseSecretRefRejectsInvalid(t *testing.T) {
	if _, _, _, err := ParseSecretRef("not-a-ref"); err == nil {
		t.Fatalf("expected error")
	}
}

func TestRedactArgsHidesAssignments(t *testing.T) {
	got := strings.Join(redactArgs([]string{"item", "edit", "item-name", "--vault", "vault", "token=secret"}), " ")
	if strings.Contains(got, "secret") {
		t.Fatalf("redacted args leaked secret: %q", got)
	}
	if !strings.Contains(got, "token=<redacted>") {
		t.Fatalf("redacted args = %q", got)
	}
}
