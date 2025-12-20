package json_parser

import (
	"encoding/json"
	"fmt"
	"os"
)

func ParseJsonConfig[T any](filePath string, config *T) error {
	configurationData, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read configuration file: %q %w", filePath, err)
	}

	err = json.Unmarshal(configurationData, config)
	if err != nil {
		return fmt.Errorf("failed to load configuration in json %q %w", filePath, err)
	}

	return nil
}
