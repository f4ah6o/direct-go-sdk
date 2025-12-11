package main

import (
	"sort"
	"time"
)

// CoverageReport represents the complete coverage analysis
type CoverageReport struct {
	Metadata   ReportMetadata   `json:"metadata"`
	Summary    CoverageSummary  `json:"summary"`
	Categories []Category       `json:"categories"`
	AllMethods MethodList       `json:"all_methods"`
}

// ReportMetadata contains report generation information
type ReportMetadata struct {
	GeneratedAt time.Time `json:"generated_at"`
	ToolVersion string    `json:"tool_version"`
	JSPath      string    `json:"js_path"`
	GoPath      string    `json:"go_path"`
}

// CoverageSummary contains high-level coverage metrics
type CoverageSummary struct {
	TotalJSMethods   int     `json:"total_js_methods"`
	TotalGoMethods   int     `json:"total_go_methods"`
	CoveragePercent  float64 `json:"coverage_percentage"`
	ImplementedCount int     `json:"implemented_count"`
	MissingCount     int     `json:"missing_count"`
}

// Category represents a functional grouping of methods with coverage info
type Category struct {
	Name             string   `json:"name"`
	TotalMethods     int      `json:"total_methods"`
	ImplementedCount int      `json:"implemented_count"`
	CoveragePercent  float64  `json:"coverage_percentage"`
	Implemented      []string `json:"implemented"`
	Missing          []string `json:"missing"`
}

// MethodList contains all methods organized by implementation status
type MethodList struct {
	JSMethods   []string `json:"js_methods"`
	GoMethods   []string `json:"go_methods"`
	Implemented []string `json:"implemented"`
	Missing     []string `json:"missing"`
}

// AnalyzeCoverage performs coverage analysis on JS and Go methods
func AnalyzeCoverage(jsMethods, goMethods []string, jsPath, goPath string) *CoverageReport {
	// Create sets for quick lookup
	jsSet := make(map[string]bool)
	for _, method := range jsMethods {
		jsSet[method] = true
	}

	goSet := make(map[string]bool)
	for _, method := range goMethods {
		goSet[method] = true
	}

	// Determine implemented and missing methods
	var implemented, missing []string
	for _, method := range jsMethods {
		if goSet[method] {
			implemented = append(implemented, method)
		} else {
			missing = append(missing, method)
		}
	}

	// Sort for consistent output
	sort.Strings(implemented)
	sort.Strings(missing)

	// Calculate overall coverage
	totalJS := len(jsMethods)
	totalGo := len(goMethods)
	implementedCount := len(implemented)
	missingCount := len(missing)
	coveragePercent := 0.0
	if totalJS > 0 {
		coveragePercent = float64(implementedCount) / float64(totalJS) * 100.0
	}

	// Analyze by category
	categories := analyzeByCategory(implemented, missing)

	// Build report
	report := &CoverageReport{
		Metadata: ReportMetadata{
			GeneratedAt: time.Now(),
			ToolVersion: "1.0.0",
			JSPath:      jsPath,
			GoPath:      goPath,
		},
		Summary: CoverageSummary{
			TotalJSMethods:   totalJS,
			TotalGoMethods:   totalGo,
			CoveragePercent:  coveragePercent,
			ImplementedCount: implementedCount,
			MissingCount:     missingCount,
		},
		Categories: categories,
		AllMethods: MethodList{
			JSMethods:   jsMethods,
			GoMethods:   goMethods,
			Implemented: implemented,
			Missing:     missing,
		},
	}

	return report
}

// analyzeByCategory breaks down coverage by functional category
func analyzeByCategory(implemented, missing []string) []Category {
	var categories []Category

	// Create sets for quick lookup
	implementedSet := make(map[string]bool)
	for _, method := range implemented {
		implementedSet[method] = true
	}

	missingSet := make(map[string]bool)
	for _, method := range missing {
		missingSet[method] = true
	}

	// Process each category in order
	for _, categoryName := range categoryOrder {
		methods := jsMethodsByCategory[categoryName]

		var categoryImplemented, categoryMissing []string

		for _, method := range methods {
			if implementedSet[method] {
				categoryImplemented = append(categoryImplemented, method)
			} else if missingSet[method] {
				categoryMissing = append(categoryMissing, method)
			}
		}

		totalMethods := len(methods)
		implementedCount := len(categoryImplemented)
		coveragePercent := 0.0
		if totalMethods > 0 {
			coveragePercent = float64(implementedCount) / float64(totalMethods) * 100.0
		}

		category := Category{
			Name:             categoryName,
			TotalMethods:     totalMethods,
			ImplementedCount: implementedCount,
			CoveragePercent:  coveragePercent,
			Implemented:      categoryImplemented,
			Missing:          categoryMissing,
		}

		categories = append(categories, category)
	}

	return categories
}

// GetCoverageStatus returns a status emoji based on coverage percentage
func GetCoverageStatus(percent float64) string {
	if percent >= 80.0 {
		return "🟢" // Green - good coverage
	} else if percent >= 50.0 {
		return "🟡" // Yellow - moderate coverage
	} else if percent >= 20.0 {
		return "🟠" // Orange - low coverage
	}
	return "🔴" // Red - very low coverage
}
