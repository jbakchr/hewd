package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"

	iscan "github.com/jbakchr/hewd/internal/scan"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// DoctorResult represents the structured diagnostics output.
type DoctorResult struct {
	MissingDocs     []string `json:"missingDocs" yaml:"missingDocs"`
	PresentDocs     []string `json:"presentDocs" yaml:"presentDocs"`
	Languages       []string `json:"languages" yaml:"languages"`
	HealthScore     int      `json:"healthScore" yaml:"healthScore"`
	Recommendations []string `json:"recommendations" yaml:"recommendations"`
}

func newDoctorCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "doctor",
		Short: "Analyze the project and report documentation and structure issues",
		Long: `Run project diagnostics based on scan results, detecting missing documentation,
project structure problems, and offering improvement suggestions.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cwd, err := os.Getwd()
			if err != nil {
				return fmt.Errorf("could not get working directory: %w", err)
			}

			// Flags for output modes
			jsonOut, _ := cmd.Flags().GetBool("json")
			yamlOut, _ := cmd.Flags().GetBool("yaml")
			pretty, _ := cmd.Flags().GetBool("pretty")

			// Prevent incompatible flag combinations
			if jsonOut && yamlOut {
				return fmt.Errorf("cannot use --json and --yaml together")
			}

			// Use your existing scanner
			summary, err := iscan.ScanDirectory(cwd)
			if err != nil {
				return err
			}

			doctor := evaluateProject(summary)

			// Structured output modes
			if jsonOut {
				return printDoctorJSON(doctor, pretty)
			}
			if yamlOut {
				return printDoctorYAML(doctor)
			}

			// Default: Human-readable report
			printDoctorReport(doctor)

			return nil
		},
	}

	cmd.Flags().Bool("json", false, "Output diagnostics in JSON format")
	cmd.Flags().Bool("yaml", false, "Output diagnostics in YAML format")
	cmd.Flags().Bool("pretty", false, "Pretty-print JSON output")

	return cmd
}

// evaluateProject converts a ScanSummary into a DoctorResult.
func evaluateProject(s *iscan.Summary) *DoctorResult {
	dr := &DoctorResult{
		MissingDocs:     []string{},
		PresentDocs:     []string{},
		Recommendations: []string{},
	}

	// Documentation health
	for doc, present := range s.Documentation {
		if present {
			dr.PresentDocs = append(dr.PresentDocs, doc)
		} else {
			dr.MissingDocs = append(dr.MissingDocs, doc)
			dr.Recommendations = append(dr.Recommendations,
				fmt.Sprintf("Add a %s file to improve project documentation.", doc))
		}
	}

	// Languages
	langs := make([]string, 0, len(s.Languages))
	for l := range s.Languages {
		langs = append(langs, l)
	}
	sort.Strings(langs)
	dr.Languages = langs

	// Simple health scoring system (0–100 scale)
	score := 100

	// Missing docs penalize
	score -= len(dr.MissingDocs) * 10

	// Too many languages? (indicates messy project)
	if len(langs) > 3 {
		score -= 10
		dr.Recommendations = append(dr.Recommendations,
			"Project uses many languages — consider organizing into submodules.")
	}

	// No languages? Strange project
	if len(langs) == 0 {
		score -= 30
		dr.Recommendations = append(dr.Recommendations,
			"No recognized source code detected.")
	}

	if score < 0 {
		score = 0
	}
	dr.HealthScore = score

	return dr
}

func printDoctorReport(dr *DoctorResult) {
	fmt.Println("Project Health Check:")

	fmt.Println()
	fmt.Println("Documentation:")
	if len(dr.PresentDocs) > 0 {
		fmt.Println("  Present:")
		for _, doc := range dr.PresentDocs {
			fmt.Println("   ✔", doc)
		}
		fmt.Println()
	}

	if len(dr.MissingDocs) > 0 {
		fmt.Println("  Missing:")
		for _, doc := range dr.MissingDocs {
			fmt.Println("   ✘", doc)
		}
	}
	fmt.Println()

	fmt.Println("Languages:")
	if len(dr.Languages) == 0 {
		fmt.Println("  (none detected)")
	} else {
		for _, l := range dr.Languages {
			fmt.Println("  -", l)
		}
	}

	fmt.Println()
	fmt.Printf("Overall Health Score: %d/100\n", dr.HealthScore)

	fmt.Println()
	fmt.Println("Recommendations:")
	if len(dr.Recommendations) == 0 {
		fmt.Println("  ✔ No issues found")
	} else {
		for _, r := range dr.Recommendations {
			fmt.Println("  -", r)
		}
	}

	fmt.Println()
	fmt.Println("Doctor complete.")
}

func printDoctorJSON(dr *DoctorResult, pretty bool) error {
	var data []byte
	var err error

	if pretty {
		data, err = json.MarshalIndent(dr, "", "  ")
	} else {
		data, err = json.Marshal(dr)
	}

	if err != nil {
		return fmt.Errorf("failed to encode JSON: %w", err)
	}

	fmt.Println(string(data))
	return nil
}

func printDoctorYAML(dr *DoctorResult) error {
	data, err := yaml.Marshal(dr)
	if err != nil {
		return fmt.Errorf("failed to encode YAML: %w", err)
	}

	fmt.Println(string(data))
	return nil
}
