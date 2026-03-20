package diff

import (
	"strings"
)

type RegressionGateResult struct {
	Failed  bool
	Reasons []string
}

func EvaluateRegressionGates(
	dr DiffResult,
	failScoreDrop int,
	failNewErrors bool,
	failAny bool,
) RegressionGateResult {

	res := RegressionGateResult{}

	// 1. Score-drop gating
	if failScoreDrop > 0 && dr.ScoreDelta <= -failScoreDrop {
		res.Failed = true
		res.Reasons = append(res.Reasons,
			"Score dropped by more than allowed threshold")
	}

	// 2. New error-level issues
	if failNewErrors {
		for _, issue := range dr.NewIssues {
			if strings.EqualFold(string(issue.Level), "error") {
				res.Failed = true
				res.Reasons = append(res.Reasons,
					"New error-level issues introduced")
				break
			}
		}
	}

	// 3. Any regression
	if failAny {
		if dr.ScoreDelta < 0 || len(dr.NewIssues) > 0 {
			res.Failed = true
			res.Reasons = append(res.Reasons,
				"General regression detected (score drop or new issues)")
		}
	}

	return res
}
