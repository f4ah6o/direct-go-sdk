// Command doccheck verifies that every Go snippet in the repository's
// README files compiles. Fenced ```go blocks are extracted from each
// README*.md (skipping vendored/source-synced directories), materialized as
// package-main programs under docs/gen/, and built inside the docs module,
// which requires the local modules via replace directives.
//
// Run from the repository root:
//
//	go run ./tools/doccheck
package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	fenceRe  = regexp.MustCompile("(?s)```(?:go|golang)\\s*\\n(.*?)```")
	skipDirs = map[string]bool{
		".git": true, ".wt": true, "issues": true, "state": true,
		"direct-js-source": true, "daab-source": true, "docs": true,
	}
)

func main() {
	root, err := os.Getwd()
	if err != nil {
		fatal(err)
	}

	docsDir := filepath.Join(root, "docs")
	genDir := filepath.Join(docsDir, "gen")
	if err := os.RemoveAll(genDir); err != nil {
		fatal(err)
	}
	if err := os.MkdirAll(genDir, 0o755); err != nil {
		fatal(err)
	}

	var readmes []string
	err = filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() && skipDirs[d.Name()] {
			return filepath.SkipDir
		}
		if !d.IsDir() && strings.HasPrefix(d.Name(), "README") && strings.HasSuffix(d.Name(), ".md") {
			readmes = append(readmes, path)
		}
		return nil
	})
	if err != nil {
		fatal(err)
	}

	count := 0
	for _, md := range readmes {
		rel, err := filepath.Rel(root, md)
		if err != nil {
			fatal(err)
		}
		data, err := os.ReadFile(md)
		if err != nil {
			fatal(err)
		}
		blocks := fenceRe.FindAllSubmatch(data, -1)
		for i, block := range blocks {
			src := string(block[1])
			if !strings.Contains(src, "package ") {
				fmt.Fprintf(os.Stderr, "::error ::%s Go snippet %d has no package clause; "+
					"every snippet must be a complete compilable program\n", rel, i)
				os.Exit(1)
			}
			slug := strings.NewReplacer("/", "_", ".", "_").Replace(strings.TrimSuffix(rel, ".md"))
			dir := filepath.Join(genDir, fmt.Sprintf("%s_%d", slug, i))
			if err := os.MkdirAll(dir, 0o755); err != nil {
				fatal(err)
			}
			if err := os.WriteFile(filepath.Join(dir, "main.go"), []byte(src), 0o644); err != nil {
				fatal(err)
			}
			count++
		}
	}
	if count == 0 {
		fmt.Fprintln(os.Stderr, "::error ::no README Go snippets found")
		os.Exit(1)
	}
	fmt.Printf("doccheck: extracted %d Go snippet(s) into %s\n", count, filepath.Join("docs", "gen"))

	build := exec.Command("go", "build", "./gen/...")
	build.Dir = docsDir
	build.Stdout = os.Stdout
	build.Stderr = os.Stderr
	if err := build.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "::error ::README Go snippets failed to compile")
		os.Exit(1)
	}
	fmt.Println("doccheck: all README Go snippets compile")
}

func fatal(err error) {
	fmt.Fprintf(os.Stderr, "doccheck: %v\n", err)
	os.Exit(1)
}
