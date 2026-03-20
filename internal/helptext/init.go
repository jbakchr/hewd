package helptext

const InitUse = "init"

const InitShort = "Initialize a new hewd configuration in the current repository."

const InitLong = `hewd init creates a .hewd/config.yaml file in the current repository. This file
customizes rule behavior, severity overrides, scoring weights, and include/exclude
paths used by all hewd commands.

The command is safe to run multiple times. Use --force to overwrite an existing
configuration.`

const InitExample = `
  # Initialize hewd config
  hewd init

  # Overwrite (force)
  hewd init --force

  # View the config
  cat .hewd/config.yaml
`
