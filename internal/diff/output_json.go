package diff

import (
	"encoding/json"
	"os"
)

func WriteJSON(out DiffOutput) error {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	return enc.Encode(out)
}
