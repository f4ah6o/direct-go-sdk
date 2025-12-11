package main

import (
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

// ExtractJSMethods extracts RPC method names from JavaScript source files
func ExtractJSMethods(jsPath string) ([]string, error) {
	methodSet := make(map[string]bool)

	// Files to check
	files := []string{
		filepath.Join(jsPath, "lib", "direct-node.js"),
		filepath.Join(jsPath, "lib", "direct.js"),
	}

	// Regex pattern to match .call("method_name"
	pattern := regexp.MustCompile(`\.call\("([a-z_]+)"`)

	for _, file := range files {
		content, err := os.ReadFile(file)
		if err != nil {
			// If file doesn't exist, try without the lib/ prefix
			file = filepath.Join(jsPath, filepath.Base(file))
			content, err = os.ReadFile(file)
			if err != nil {
				continue // Skip files that don't exist
			}
		}

		// Find all matches
		matches := pattern.FindAllSubmatch(content, -1)
		for _, match := range matches {
			if len(match) > 1 {
				method := string(match[1])
				methodSet[method] = true
			}
		}
	}

	// Convert set to sorted slice
	methods := make([]string, 0, len(methodSet))
	for method := range methodSet {
		methods = append(methods, method)
	}
	sort.Strings(methods)

	return methods, nil
}

// ExtractGoMethods extracts RPC method names from Go source files
func ExtractGoMethods(goPath string) ([]string, error) {
	methodSet := make(map[string]bool)

	// Regex patterns to match c.call("method_name") and c.Call("method_name")
	patterns := []*regexp.Regexp{
		regexp.MustCompile(`c\.call\("([a-z_]+)"`),
		regexp.MustCompile(`c\.Call\("([a-z_]+)"`),
	}

	// Walk through all .go files in the directory
	err := filepath.Walk(goPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories and non-.go files
		if info.IsDir() || !strings.HasSuffix(path, ".go") {
			return nil
		}

		// Skip test files
		if strings.HasSuffix(path, "_test.go") {
			return nil
		}

		// Skip tools directory (avoid self-reference)
		if strings.Contains(path, string(filepath.Separator)+"tools"+string(filepath.Separator)) {
			return nil
		}

		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		// Try all patterns
		for _, pattern := range patterns {
			matches := pattern.FindAllSubmatch(content, -1)
			for _, match := range matches {
				if len(match) > 1 {
					method := string(match[1])
					methodSet[method] = true
				}
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	// Convert set to sorted slice
	methods := make([]string, 0, len(methodSet))
	for method := range methodSet {
		methods = append(methods, method)
	}
	sort.Strings(methods)

	return methods, nil
}

// ValidateExtraction compares extracted methods with baseline
func ValidateExtraction(extracted []string, baseline []string, source string) {
	if len(extracted) == 0 {
		return
	}

	extractedSet := make(map[string]bool)
	for _, method := range extracted {
		extractedSet[method] = true
	}

	baselineSet := make(map[string]bool)
	for _, method := range baseline {
		baselineSet[method] = true
	}

	// Find methods in extracted but not in baseline (new methods)
	var newMethods []string
	for method := range extractedSet {
		if !baselineSet[method] {
			newMethods = append(newMethods, method)
		}
	}

	// Find methods in baseline but not in extracted (missing methods)
	var missingMethods []string
	for method := range baselineSet {
		if !extractedSet[method] {
			missingMethods = append(missingMethods, method)
		}
	}

	if len(newMethods) > 0 {
		sort.Strings(newMethods)
	}

	if len(missingMethods) > 0 {
		sort.Strings(missingMethods)
	}
}
