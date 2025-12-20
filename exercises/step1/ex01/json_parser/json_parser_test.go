package json_parser

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

type SampleConfiguration struct {
	Name string `json:"name"`
}

func getArtifactFullPath(relativePath string) string {
	_, thisFile, _, _ := runtime.Caller(0)
	dir := filepath.Dir(thisFile)

	return filepath.Join(dir, "testdata", relativePath)
}

func TestParseJsonConfigNonExistingFile(t *testing.T) {
	var config SampleConfiguration
	err := ParseJsonConfig("some_missing_file.json", &config)
	if !errors.Is(err, os.ErrNotExist) {
		t.Fatal("Parse JSON should file on missing file with os.ErrNotExist error")
	}
}

func TestParseJsonConfigLoadIntoMap(t *testing.T) {
	config := make(map[string]any)
	err := ParseJsonConfig(getArtifactFullPath("sample.json"), &config)
	if err != nil {
		t.Fatalf("Parse JSON failed with error %v", err)
	}
	expectedValue := "test"
	if fieldValue, ok := config["name"]; !ok || fieldValue != expectedValue {
		t.Fatalf("Got %q wanted %q", fieldValue, expectedValue)
	}
}

func TestParseJsonConfigLoadIntoStruct(t *testing.T) {
	var config SampleConfiguration
	err := ParseJsonConfig(getArtifactFullPath("sample.json"), &config)
	if err != nil {
		t.Fatalf("Parse JSON failed with error %v", err)
	}
	if expectedField := "test"; config.Name != expectedField {
		t.Fatalf("Got %q wanted %q", config.Name, expectedField)
	}
}

func TestParseJsonConfigInvalidJSON(t *testing.T) {
	var config SampleConfiguration

	err := ParseJsonConfig(getArtifactFullPath("invalid_json.json"), &config)
	if err == nil {
		t.Fatal("expected error for invalid JSON, got nil")
	}

	// Demonstrate errors.As() - verify the underlying error type is preserved
	var jsonErr *json.SyntaxError
	if !errors.As(err, &jsonErr) {
		t.Fatalf("expected json.SyntaxError in error chain, got: %T", err)
	}

	t.Logf("Successfully detected JSON syntax error: %v", jsonErr)
}
