package debuglog

import (
	"bytes"
	"strconv"
	"strings"
	"testing"
)

func TestSummarizePayloadDoesNotIncludeValues(t *testing.T) {
	secret := "access-token-secret"
	message := "private message body"
	authorization := "Bearer authorization-secret"
	payload := map[string]interface{}{
		"access_token":  secret,
		"text":          message,
		"Authorization": authorization,
	}

	got := SummarizePayload(payload)
	for _, unwanted := range []string{secret, message, authorization} {
		if strings.Contains(got, unwanted) {
			t.Fatalf("SummarizePayload leaked %q in %q", unwanted, got)
		}
	}
	if got != "map(len=3)" {
		t.Fatalf("SummarizePayload() = %q, want map metadata", got)
	}
}

func TestRedactionHelpers(t *testing.T) {
	for _, test := range []struct {
		name string
		got  string
		want string
	}{
		{name: "id", got: RedactID("user-123"), want: "id:"},
		{name: "secret", got: RedactSecret("secret"), want: "<redacted secret>"},
		{name: "authorization", got: RedactAuthorization("Bearer secret"), want: "<redacted authorization>"},
	} {
		t.Run(test.name, func(t *testing.T) {
			if !strings.HasPrefix(test.got, test.want) && test.got != test.want {
				t.Fatalf("got %q, want prefix/value %q", test.got, test.want)
			}
		})
	}
	if got := RedactID("user-123"); got == "user-123" || got != RedactID("user-123") {
		t.Fatalf("RedactID() is not stable and non-reversible: %q", got)
	}
}

func TestUnsafePayloadTracingRequiresExplicitOptIn(t *testing.T) {
	var output bytes.Buffer
	logger := NewLogger(LoggerOptions{Level: LevelVerbose, Writer: &output})
	logger.UnsafePrintf("payload=%s", "private message body")
	if strings.Contains(output.String(), "private message body") {
		t.Fatalf("unsafe payload was logged without opt-in: %q", output.String())
	}

	logger.EnableUnsafePayloadTracing()
	logger.UnsafePrintf("payload=%s", "private message body")
	if !strings.Contains(output.String(), "private message body") {
		t.Fatalf("explicit unsafe payload was not logged: %q", output.String())
	}
	if !strings.Contains(output.String(), "UNSAFE payload tracing enabled") {
		t.Fatalf("unsafe warning missing: %q", output.String())
	}
}

func TestLoggerInstancesAreIndependent(t *testing.T) {
	var firstOutput, secondOutput bytes.Buffer
	first := NewLogger(LoggerOptions{Level: LevelNormal, Writer: &firstOutput})
	second := NewLogger(LoggerOptions{Level: LevelOff, Writer: &secondOutput})

	first.Printf("method=%s", "get_me")
	second.Printf("method=%s", "get_me")

	if !strings.Contains(firstOutput.String(), "get_me") {
		t.Fatalf("enabled logger did not emit diagnostics: %q", firstOutput.String())
	}
	if secondOutput.Len() != 0 {
		t.Fatalf("disabled logger emitted diagnostics: %q", secondOutput.String())
	}
}

func TestSafeDiagnosticsAtNormalAndVerboseLevels(t *testing.T) {
	const (
		accessToken = "access-token-secret"
		messageBody = "private message body"
	)
	payload := map[string]interface{}{
		"access_token":  accessToken,
		"content":       messageBody,
		"Authorization": "Bearer authorization-secret",
	}

	for _, level := range []int{LevelNormal, LevelVerbose} {
		t.Run(strconv.Itoa(level), func(t *testing.T) {
			var output bytes.Buffer
			logger := NewLogger(LoggerOptions{Level: level, Writer: &output})
			logger.Printf("session token=%s message=%s payload=%s", RedactSecret(accessToken), SummarizePayload(messageBody), SummarizePayload(payload))
			logger.Verbose("verbose payload=%s", SummarizePayload(payload))

			got := output.String()
			for _, unwanted := range []string{accessToken, messageBody, "authorization-secret", "Bearer"} {
				if strings.Contains(got, unwanted) {
					t.Fatalf("safe diagnostics leaked %q: %q", unwanted, got)
				}
			}
			if !strings.Contains(got, "<redacted secret>") || !strings.Contains(got, "string(len=") {
				t.Fatalf("safe diagnostics lost expected metadata: %q", got)
			}
		})
	}
}

func TestWriterDoesNotLogPayloadWithoutUnsafeOptIn(t *testing.T) {
	const messageBody = "private stream body"
	var output bytes.Buffer
	logger := NewLogger(LoggerOptions{Level: LevelNormal, Writer: &output})

	if _, err := logger.Writer().Write([]byte(messageBody)); err != nil {
		t.Fatalf("safe writer failed: %v", err)
	}
	if strings.Contains(output.String(), messageBody) {
		t.Fatalf("safe writer leaked payload: %q", output.String())
	}
	if !strings.Contains(output.String(), "writer bytes=") {
		t.Fatalf("safe writer did not record byte count: %q", output.String())
	}

	logger.EnableUnsafePayloadTracing()
	if _, err := logger.UnsafeWriter().Write([]byte(messageBody)); err != nil {
		t.Fatalf("unsafe writer failed: %v", err)
	}
	if !strings.Contains(output.String(), messageBody) {
		t.Fatalf("explicit unsafe writer did not emit payload: %q", output.String())
	}
}
