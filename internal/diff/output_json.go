package diff

import (
	"encoding/json"
	"os"

	"github.com/jbakchr/hewd/internal/cliutils"
)

func WriteJSON(out DiffOutput) error {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")

	if err := enc.Encode(out); err != nil {
		return cliutils.ErrHint(
			"failed to encode json output",
			"try running with --pretty disabled or check for invalid characters in your report",
		)
	}

	return nil
}
