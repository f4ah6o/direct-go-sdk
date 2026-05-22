package main

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestGenerateBadge(t *testing.T) {
	report := AnalyzeCoverage(
		[]string{"get_me", "get_users"},
		[]string{"get_me"},
		"/tmp/js",
		"/tmp/go",
	)

	data, err := GenerateBadge(report)
	if err != nil {
		t.Fatalf("GenerateBadge returned error: %v", err)
	}

	var badge Badge
	if err := json.Unmarshal(data, &badge); err != nil {
		t.Fatalf("badge is invalid JSON: %v", err)
	}
	if badge.SchemaVersion != 1 {
		t.Fatalf("SchemaVersion = %d, want 1", badge.SchemaVersion)
	}
	if badge.Label != "direct-go porting" {
		t.Fatalf("Label = %q", badge.Label)
	}
	if badge.Message != "50.0% (1/2)" {
		t.Fatalf("Message = %q", badge.Message)
	}
	if badge.Color != "yellow" {
		t.Fatalf("Color = %q", badge.Color)
	}
}

func TestGenerateMarkdownIsStableAndShowsProgress(t *testing.T) {
	report := AnalyzeCoverage(
		[]string{"get_me", "get_users"},
		[]string{"get_me"},
		"/tmp/js",
		"/tmp/go",
	)

	markdown := GenerateMarkdown(report)
	if strings.Contains(markdown, "Generated**:") {
		t.Fatalf("markdown should not include a timestamp:\n%s", markdown)
	}
	if !strings.Contains(markdown, "`###############---------------` 50.00%") {
		t.Fatalf("markdown does not include summary progress bar:\n%s", markdown)
	}
	if !strings.Contains(markdown, "| User Management | 1/11 | 9.1% | `#-----------` |") {
		t.Fatalf("markdown does not include category progress:\n%s", markdown)
	}
}
