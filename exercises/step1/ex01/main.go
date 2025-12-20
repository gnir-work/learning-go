//Build a simple configuration parser that demonstrates idiomatic error handling.
//
//Task: Create a package that reads a JSON config file and validates required fields.
//
//Implement custom error types using errors.New() and fmt.Errorf() with %w
//Use error wrapping and errors.Is() / errors.As()
//Return sentinel errors for specific cases (e.g., ErrMissingField)
//Write a CLI tool that uses this package and handles errors gracefully
//Focus: Error wrapping, sentinel errors, error context

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gnir-work/learning-go/exercises/step1/ex01/json_parser"
)

func loadConfigurationFilePathFromCLI() string {
	var configurationFilePath string
	flag.StringVar(&configurationFilePath, "config", "", "Path to configuration file")
	flag.Parse()

	if configurationFilePath == "" {
		fmt.Fprintf(os.Stderr, "missing required -config\n")
		flag.Usage()
		os.Exit(2)
	}
	return configurationFilePath
}

func main() {
	configurationFilePath := loadConfigurationFilePathFromCLI()

	configuration := make(map[string]any)
	err := json_parser.ParseJsonConfig(configurationFilePath, &configuration)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load configuration: '%v'", err)
		os.Exit(1)
	}
	fmt.Printf("Loaded configuration: '%v'\n", configuration)

}
