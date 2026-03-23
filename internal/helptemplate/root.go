package helptemplate

import (
	"github.com/jbakchr/hewd/internal/formatter"
)

var RootHelpTemplate = `{{with .Long}}{{.}}{{else}}{{.Short}}{{end}}` +

	formatter.WhiteBold("Usage:") + `
  
  {{if .Runnable}}{{.UseLine}}{{end}}{{if .HasAvailableSubCommands}}{{.CommandPath}} [command] [flags]{{end}}` +

	formatter.WhiteBold("\n\nExamples:") + `
	{{if .HasExample}}{{.Example}}{{end}}` +

	formatter.WhiteBold("\nAnalysis Commands:") + `

{{range .Commands}}{{if eq .GroupID "analysis"}}  {{rpad .Name .NamePadding }} {{.Short}}
{{end}}{{end}}` +
	formatter.WhiteBold("\nMaintenance Commands:") + `

{{range .Commands}}{{if eq .GroupID "maintenance"}}  {{rpad .Name .NamePadding }} {{.Short}}
{{end}}{{end}}` +
	formatter.WhiteBold("\nReporting Commands:") + `

{{range .Commands}}{{if eq .GroupID "reporting"}}  {{rpad .Name .NamePadding }} {{.Short}}
{{end}}{{end}}` +

	formatter.WhiteBold("\nAdditional Commands:") + `

{{range .Commands}}{{if not (or (eq .GroupID "analysis") (eq .GroupID "maintenance") (eq .GroupID "reporting"))}}  {{rpad .Name .NamePadding }} {{.Short}}.
{{end}}{{end}}` +
	formatter.WhiteBold("\nFlags:") + `

{{.Flags.FlagUsages}}
{{if .HasAvailableSubCommands}}{{.CommandPath}} [command] --help for more information about a command.{{end}}
`
