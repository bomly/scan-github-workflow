package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"strings"
)

var validScopes = []string{"require", "tool"}

type Dependency struct {
	Module  string `json:"module"`
	Version string `json:"version"`
	Scope   string `json:"scope"`
	Type    string `json:"type"`
}

type Module struct {
	Name         string       `json:"name"`
	File         string       `json:"file"`
	Version      string       `json:"version"`
	Dependencies []Dependency `json:"dependencies"`
}

type ScanResult struct {
	Version string   `json:"version"`
	System  string   `json:"system"`
	Scopes  []string `json:"scopes"`
	Modules []Module `json:"modules"`
}

func main() {
	err := runScanner()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}

func runScanner() error {
	scopes, err := getScopes()
	if err != nil {
		return err
	}

	result, err := scan(scopes)
	if err != nil {
		return err
	}

	err = printResult(result)
	if err != nil {
		return err
	}

	return nil
}

func printResult(result ScanResult) error {
	var buf bytes.Buffer
	var enc = json.NewEncoder(&buf)
	enc.SetIndent("", "  ")

	err := enc.Encode(result)
	if err != nil {
		return fmt.Errorf("failed to encode dependencies: %w", err)
	}

	fmt.Println(buf.String())

	return nil
}

func scan(scopes []string) (ScanResult, error) {
	return ScanResult{
		Version: "1.0.0",
		System:  "github-workflow",
		Scopes:  scopes,
		Modules: []Module{
			{
				Name:    "Go",
				File:    ".github/workflows/go.yaml",
				Version: "git sha",
				Dependencies: []Dependency{
					{
						Module:  "actions/checkout",
						Version: "v4",
						Scope:   "build",
						Type:    "github-workflow",
					},
					{
						Module:  "actions/setup-g",
						Version: "v5",
						Scope:   "build",
						Type:    "github-workflow",
					},
					{
						Module:  "codecov/codecov-action",
						Version: "v5",
						Scope:   "build",
						Type:    "github-workflow",
					},
				},
			},
		},
	}, nil
}

func getScopes() ([]string, error) {
	scopeList := strings.TrimSpace(os.Getenv("SCOPES"))
	var scopes []string
	for _, s := range strings.Split(scopeList, ",") {
		scope := strings.TrimSpace(s)
		if scope == "" {
			continue
		}

		if !slices.Contains(validScopes, scope) {
			return nil, fmt.Errorf("invalid scope %q, valid scopes are %v", scope, validScopes)
		}

		scopes = append(scopes, strings.TrimSpace(scope))
	}

	if len(scopes) == 0 {
		return []string{"require"}, nil
	}

	return scopes, nil
}
