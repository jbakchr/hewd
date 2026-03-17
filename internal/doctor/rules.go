package doctor

import "github.com/jbakchr/hewd/internal/scan"

// Built-in diagnostic rules
var Rules = []Rule{
    {
        ID:          "DOC-001",
        Description: "Project should have a README.md",
        Check: func(s *scan.Summary) (bool, string) {
            if s.Documentation["README.md"] {
                return true, "README.md present"
            }
            return false, "README.md missing"
        },
    },
    {
        ID:          "DOC-002",
        Description: "Project should have a LICENSE file",
        Check: func(s *scan.Summary) (bool, string) {
            if s.Documentation["LICENSE"] {
                return true, "LICENSE present"
            }
            return false, "LICENSE missing"
        },
    },
    {
        ID:          "CFG-001",
        Description: "Project should have at least one configuration file",
        Check: func(s *scan.Summary) (bool, string) {
            if len(s.ConfigFiles) > 0 {
                return true, "Configuration files detected"
            }
            return false, "No configuration files detected"
        },
    },
}