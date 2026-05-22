package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
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
	constants := make(map[string]string)
	files := make([]*ast.File, 0)

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

		fileSet := token.NewFileSet()
		file, err := parser.ParseFile(fileSet, path, nil, 0)
		if err != nil {
			return err
		}

		files = append(files, file)

		return nil
	})

	if err != nil {
		return nil, err
	}

	for _, file := range files {
		collectStringConstants(file, constants)
	}
	for _, file := range files {
		collectRPCCalls(file, constants, methodSet)
	}

	// Convert set to sorted slice
	methods := make([]string, 0, len(methodSet))
	for method := range methodSet {
		methods = append(methods, method)
	}
	sort.Strings(methods)

	return methods, nil
}

func collectRPCCalls(file *ast.File, constants map[string]string, methodSet map[string]bool) {
	ast.Inspect(file, func(node ast.Node) bool {
		call, ok := node.(*ast.CallExpr)
		if !ok || len(call.Args) == 0 {
			return true
		}
		if !isRPCCall(call) {
			return true
		}
		if method, ok := rpcMethodName(call.Args[0], constants); ok {
			methodSet[method] = true
		}
		return true
	})
}

func collectStringConstants(file *ast.File, constants map[string]string) {
	for _, decl := range file.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok || genDecl.Tok != token.CONST {
			continue
		}
		for _, spec := range genDecl.Specs {
			valueSpec, ok := spec.(*ast.ValueSpec)
			if !ok {
				continue
			}
			for i, name := range valueSpec.Names {
				if !strings.HasPrefix(name.Name, "Method") || i >= len(valueSpec.Values) {
					continue
				}
				if value, ok := stringLiteral(valueSpec.Values[i]); ok {
					constants[name.Name] = value
				}
			}
		}
	}
}

func isRPCCall(call *ast.CallExpr) bool {
	selector, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}
	return selector.Sel.Name == "Call" || selector.Sel.Name == "call"
}

func rpcMethodName(expr ast.Expr, constants map[string]string) (string, bool) {
	if value, ok := stringLiteral(expr); ok {
		return value, true
	}
	if ident, ok := expr.(*ast.Ident); ok {
		value, found := constants[ident.Name]
		return value, found
	}
	return "", false
}

func stringLiteral(expr ast.Expr) (string, bool) {
	lit, ok := expr.(*ast.BasicLit)
	if !ok || lit.Kind != token.STRING {
		return "", false
	}
	value, err := strconv.Unquote(lit.Value)
	if err != nil {
		return "", false
	}
	if value == "" {
		return "", false
	}
	return value, true
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
