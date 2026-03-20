package diff

import (
	"os"

	"gopkg.in/yaml.v3"
)

func WriteYAML(out DiffOutput) error {
	enc := yaml.NewEncoder(os.Stdout)
	enc.SetIndent(2)
	return enc.Encode(out)
}
