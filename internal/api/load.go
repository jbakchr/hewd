package api

import (
	"encoding/json"
	"fmt"
	"os"
)

func LoadReport(path string) (MachineOutput, error) {
	var m MachineOutput

	data, err := os.ReadFile(path)
	if err != nil {
		return m, err
	}

	if err := json.Unmarshal(data, &m); err != nil {
		return m, err
	}

	if m.SchemaVersion != 1 {
		return m, fmt.Errorf("unsupported schema version: %d", m.SchemaVersion)
	}

	return m, nil
}
