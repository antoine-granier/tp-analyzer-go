package reporter

import (
	"encoding/json"
	"main/internal/analyzer"
	"os"
)

func Export(results []analyzer.Result, path string) error {
	data, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}
