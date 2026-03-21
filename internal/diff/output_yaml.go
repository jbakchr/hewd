package diff

import (
	"os"

	"gopkg.in/yaml.v3"

	"github.com/jbakchr/hewd/internal/cliutils"
)

func WriteYAML(out DiffOutput) error {
	enc := yaml.NewEncoder(os.Stdout)
	enc.SetIndent(2)

	if err := enc.Encode(out); err != nil {
		enc.Close() // best effort
		return cliutils.ErrHint(
			"failed to encode yaml output",
			"try running with --pretty disabled or check for invalid characters in your report",
		)
	}

	if err := enc.Close(); err != nil {
		return cliutils.ErrHint(
			"failed to finalize yaml output",
			"ensure your output stream is writable",
		)
	}

	return nil
}
