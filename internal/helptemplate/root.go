package helptemplate

import "github.com/jbakchr/hewd/internal/cliutils"

const RootHelpTemplate = `{{with .Long}}{{.}}{{else}}{{.Short}}{{end}}` +

	cliutils.WhiteBold + "Usage:" + cliutils.Reset + `
  
  {{if .Runnable}}{{.UseLine}}{{end}}{{if .HasAvailableSubCommands}}{{.CommandPath}} [command] [flags]{{end}}` +

	cliutils.Bold + "\n\nExamples: " + cliutils.Reset + `
	{{if .HasExample}}{{.Example}}{{end}}` +

	cliutils.WhiteBold + "\nAnalysis Commands:" + cliutils.Reset + `

{{range .Commands}}{{if eq .GroupID "analysis"}}  {{rpad .Name .NamePadding }} {{.Short}}
{{end}}{{end}}` +
	cliutils.WhiteBold + "\nMaintenance Commands:" + cliutils.Reset + `

{{range .Commands}}{{if eq .GroupID "maintenance"}}  {{rpad .Name .NamePadding }} {{.Short}}
{{end}}{{end}}` +
	cliutils.WhiteBold + "\nReporting Commands:" + cliutils.Reset + `

{{range .Commands}}{{if eq .GroupID "reporting"}}  {{rpad .Name .NamePadding }} {{.Short}}
{{end}}{{end}}` +

	cliutils.Bold + "\nAdditional Commands:" + cliutils.Reset + `

{{range .Commands}}{{if not (or (eq .GroupID "analysis") (eq .GroupID "maintenance") (eq .GroupID "reporting"))}}  {{rpad .Name .NamePadding }} {{.Short}}.
{{end}}{{end}}` +
	cliutils.Bold + "\nFlags:" + cliutils.Reset + `

{{.Flags.FlagUsages}}
{{if .HasAvailableSubCommands}}{{.CommandPath}} [command] --help for more information about a command.{{end}}
`
