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
	jsPath := flag.String("js-path", "../direct-js", "Path to direct-js directory")
	goPath := flag.String("go-path", "../..", "Path to direct-go directory")
	output := flag.String("output", "", "Output file path (default: stdout)")
	format := flag.String("format", "markdown", "Output format: json|markdown|text")
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

	// Convert to absolute paths
	jsPathAbs, err := filepath.Abs(*jsPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error resolving JS path: %v\n", err)
		os.Exit(1)
	}

	goPathAbs, err := filepath.Abs(*goPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error resolving Go path: %v\n", err)
		os.Exit(1)
	}

	if *verbose {
		fmt.Fprintf(os.Stderr, "JS Path: %s\n", jsPathAbs)
		fmt.Fprintf(os.Stderr, "Go Path: %s\n", goPathAbs)
		fmt.Fprintf(os.Stderr, "Output Format: %s\n", *format)
		if *output != "" {
			fmt.Fprintf(os.Stderr, "Output File: %s\n", *output)
		}
		fmt.Fprintf(os.Stderr, "\n")
	}

	// Step 1: Extract JS methods
	var jsMethods []string
	if *useBaseline {
		if *verbose {
			fmt.Fprintf(os.Stderr, "Using hardcoded baseline for JS methods\n")
		}
		jsMethods = getAllJSMethods()
	} else {
		if *verbose {
			fmt.Fprintf(os.Stderr, "Extracting JS methods from source...\n")
		}
		jsMethods, err = ExtractJSMethods(jsPathAbs)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error extracting JS methods: %v\n", err)
			os.Exit(1)
		}
		if *verbose {
			fmt.Fprintf(os.Stderr, "Found %d JS methods\n", len(jsMethods))
		}
	}

	// Step 2: Extract Go methods
	if *verbose {
		fmt.Fprintf(os.Stderr, "Extracting Go methods from source...\n")
	}
	goMethods, err := ExtractGoMethods(goPathAbs)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error extracting Go methods: %v\n", err)
		os.Exit(1)
	}
	if *verbose {
		fmt.Fprintf(os.Stderr, "Found %d Go methods\n", len(goMethods))
	}

	// Step 3: Validate extraction (optional)
	if *verbose && !*useBaseline {
		baselineMethods := getAllJSMethods()
		ValidateExtraction(jsMethods, baselineMethods, "JavaScript")
	}

	// Step 4: Analyze coverage
	if *verbose {
		fmt.Fprintf(os.Stderr, "Analyzing coverage...\n")
	}
	report := AnalyzeCoverage(jsMethods, goMethods, jsPathAbs, goPathAbs)

	// Step 5: Generate output
	var outputContent string
	var outputBytes []byte

	switch *format {
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

	default:
		fmt.Fprintf(os.Stderr, "Unknown format: %s (use json, markdown, or text)\n", *format)
		os.Exit(1)
	}

	// Step 6: Write output
	if *output == "" {
		// Write to stdout
		fmt.Print(outputContent)
	} else {
		// Write to file
		err = os.WriteFile(*output, []byte(outputContent), 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error writing to file: %v\n", err)
			os.Exit(1)
		}
		if *verbose {
			fmt.Fprintf(os.Stderr, "Report written to: %s\n", *output)
		} else {
			fmt.Printf("Coverage report written to: %s\n", *output)
		}
	}

	// Show summary to stderr if writing to file
	if *output != "" && !*verbose {
		summary := GenerateTextSummary(report)
		fmt.Fprint(os.Stderr, "\n"+summary+"\n")
	}
}
