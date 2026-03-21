package helptext

const InitUse = "init"

const InitShort = "Initialize a new hewd configuration in the current repository."

const InitLong = `hewd init creates a .hewd/config.yaml file in the current repository. This file
allows customization of rule behavior, severity levels, scoring weights, and
scan include/exclude paths.

Features:

  • Creates .hewd directory if missing
  • Writes default configuration file
  • Safe by default (never overwrites)
  • Supports --force to replace existing config

Use hewd init to standardize configuration across repositories or to prepare a
project for diagnostics and CI usage.`

const InitExample = `
  # Initialize configuration
  hewd init

  # Overwrite existing config
  hewd init --force

  # View generated configuration
  cat .hewd/config.yaml
`
