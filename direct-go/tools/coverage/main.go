package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

const version = "1.0.0"

func main() {
	// Command-line flags
	jsPath := flag.String("js-path", "", "Path to direct-js source directory (default: auto-detect)")
	goPath := flag.String("go-path", "", "Path to direct-go directory (default: auto-detect)")
	output := flag.String("output", "", "Output file path (default: stdout)")
	format := flag.String("format", "markdown", "Output format: json|markdown|text|badge")
	verbose := flag.Bool("verbose", false, "Verbose output with extraction details")
	showVersion := flag.Bool("version", false, "Show version information")
	useBaseline := flag.Bool("use-baseline", false, "Use hardcoded baseline instead of extracting from JS")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Direct4B Porting Coverage Tool v%s\n\n", version)
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nExamples:\n")
		fmt.Fprintf(os.Stderr, "  # Generate Markdown report to stdout\n")
		fmt.Fprintf(os.Stderr, "  %s\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  # Generate JSON report to file\n")
		fmt.Fprintf(os.Stderr, "  %s -format json -output coverage.json\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  # Use specific paths\n")
		fmt.Fprintf(os.Stderr, "  %s -js-path ../direct-js -go-path .\n\n", os.Args[0])
	}

	flag.Parse()

	// Show version and exit
	if *showVersion {
		fmt.Printf("direct4b-coverage-tool v%s\n", version)
		os.Exit(0)
	}

	paths, err := resolvePaths(*jsPath, *goPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error resolving paths: %v\n", err)
		os.Exit(1)
	}
	jsPathAbs := paths.jsPath
	goPathAbs := paths.goPath

	// Convert output path to an absolute path only after cwd-independent inputs
	// have been resolved.
	outputPath := *output
	if outputPath != "" {
		outputPath, err = filepath.Abs(outputPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error resolving output path: %v\n", err)
			os.Exit(1)
		}
	}

	if *verbose {
		fmt.Fprintf(os.Stderr, "JS Path: %s\n", jsPathAbs)
		fmt.Fprintf(os.Stderr, "Go Path: %s\n", goPathAbs)
		fmt.Fprintf(os.Stderr, "Output Format: %s\n", *format)
		if outputPath != "" {
			fmt.Fprintf(os.Stderr, "Output File: %s\n", outputPath)
		}
		fmt.Fprintf(os.Stderr, "\n")
	}
	run(jsPathAbs, goPathAbs, outputPath, *format, *verbose, *useBaseline)
}

type resolvedPaths struct {
	jsPath string
	goPath string
}

func resolvePaths(jsPath, goPath string) (resolvedPaths, error) {
	if goPath == "" {
		detected, err := detectDirectGoRoot()
		if err != nil {
			return resolvedPaths{}, err
		}
		goPath = detected
	}

	goPathAbs, err := filepath.Abs(goPath)
	if err != nil {
		return resolvedPaths{}, fmt.Errorf("go path: %w", err)
	}

	if jsPath == "" {
		jsPath = filepath.Join(goPathAbs, "direct-js-source")
	}

	jsPathAbs, err := filepath.Abs(jsPath)
	if err != nil {
		return resolvedPaths{}, fmt.Errorf("js path: %w", err)
	}

	return resolvedPaths{jsPath: jsPathAbs, goPath: goPathAbs}, nil
}

func detectDirectGoRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for {
		if fileExists(filepath.Join(dir, "go.mod")) && fileExists(filepath.Join(dir, "client.go")) && fileExists(filepath.Join(dir, "events.go")) {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("could not find direct-go root from %s", dir)
		}
		dir = parent
	}
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func run(jsPathAbs, goPathAbs, output, format string, verbose, useBaseline bool) {
	// Step 1: Extract JS methods
	var jsMethods []string
	var err error
	if useBaseline {
		if verbose {
			fmt.Fprintf(os.Stderr, "Using hardcoded baseline for JS methods\n")
		}
		jsMethods = getAllJSMethods()
	} else {
		if verbose {
			fmt.Fprintf(os.Stderr, "Extracting JS methods from source...\n")
		}
		jsMethods, err = ExtractJSMethods(jsPathAbs)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error extracting JS methods: %v\n", err)
			os.Exit(1)
		}
		if verbose {
			fmt.Fprintf(os.Stderr, "Found %d JS methods\n", len(jsMethods))
		}
	}
	if len(jsMethods) == 0 {
		fmt.Fprintf(os.Stderr, "Error extracting JS methods: no methods found in %s\n", jsPathAbs)
		os.Exit(1)
	}

	// Step 2: Extract Go methods
	if verbose {
		fmt.Fprintf(os.Stderr, "Extracting Go methods from source...\n")
	}
	goMethods, err := ExtractGoMethods(goPathAbs)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error extracting Go methods: %v\n", err)
		os.Exit(1)
	}
	if verbose {
		fmt.Fprintf(os.Stderr, "Found %d Go methods\n", len(goMethods))
	}

	// Step 3: Validate extraction (optional)
	if verbose && !useBaseline {
		baselineMethods := getAllJSMethods()
		ValidateExtraction(jsMethods, baselineMethods, "JavaScript")
	}

	// Step 4: Analyze coverage
	if verbose {
		fmt.Fprintf(os.Stderr, "Analyzing coverage...\n")
	}
	report := AnalyzeCoverage(jsMethods, goMethods, jsPathAbs, goPathAbs)

	// Step 5: Generate output
	var outputContent string
	var outputBytes []byte

	switch format {
	case "json":
		outputBytes, err = GenerateJSON(report)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error generating JSON: %v\n", err)
			os.Exit(1)
		}
		outputContent = string(outputBytes)

	case "markdown", "md":
		outputContent = GenerateMarkdown(report)

	case "text", "txt":
		outputContent = GenerateTextSummary(report)

	case "badge":
		outputBytes, err = GenerateBadge(report)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error generating badge: %v\n", err)
			os.Exit(1)
		}
		outputContent = string(outputBytes)

	default:
		fmt.Fprintf(os.Stderr, "Unknown format: %s (use json, markdown, text, or badge)\n", format)
		os.Exit(1)
	}

	// Step 6: Write output
	if output == "" {
		// Write to stdout
		fmt.Print(outputContent)
	} else {
		// Write to file
		err = os.MkdirAll(filepath.Dir(output), 0755)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating output directory: %v\n", err)
			os.Exit(1)
		}
		err = os.WriteFile(output, []byte(outputContent), 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error writing to file: %v\n", err)
			os.Exit(1)
		}
		if verbose {
			fmt.Fprintf(os.Stderr, "Report written to: %s\n", output)
		} else {
			fmt.Printf("Coverage report written to: %s\n", output)
		}
	}

	// Show summary to stderr if writing to file
	if output != "" && !verbose {
		summary := GenerateTextSummary(report)
		fmt.Fprint(os.Stderr, "\n"+summary+"\n")
	}
}
