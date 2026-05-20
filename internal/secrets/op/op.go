package op

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"strings"
)

type Runner struct {
	Binary string
}

func (r Runner) bin() string {
	if r.Binary == "" {
		return "op"
	}
	return r.Binary
}

func (r Runner) Read(ctx context.Context, ref string) (string, error) {
	out, err := r.command(ctx, "read", ref)
	if err != nil {
		return "", err
	}
	return strings.TrimRight(string(out), "\r\n"), nil
}

func (r Runner) Write(ctx context.Context, ref, value string) error {
	vault, item, field, err := ParseSecretRef(ref)
	if err != nil {
		return err
	}
	assignment := field + "=" + value
	_, err = r.command(ctx, "item", "edit", item, "--vault", vault, assignment)
	return err
}

func (r Runner) command(ctx context.Context, args ...string) ([]byte, error) {
	cmd := exec.CommandContext(ctx, r.bin(), args...)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("%s %s failed: %w: %s", r.bin(), strings.Join(args, " "), err, strings.TrimSpace(stderr.String()))
	}
	return out, nil
}

func ParseSecretRef(ref string) (vault, item, field string, err error) {
	const prefix = "op://"
	if !strings.HasPrefix(ref, prefix) {
		return "", "", "", fmt.Errorf("secret reference must start with %q", prefix)
	}
	parts := strings.Split(strings.TrimPrefix(ref, prefix), "/")
	if len(parts) < 3 {
		return "", "", "", fmt.Errorf("secret reference must be op://vault/item/field")
	}
	vault = parts[0]
	item = parts[1]
	field = strings.Join(parts[2:], "/")
	if vault == "" || item == "" || field == "" {
		return "", "", "", fmt.Errorf("secret reference must include vault, item, and field")
	}
	return vault, item, field, nil
}
