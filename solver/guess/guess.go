package guess

import (
	"github.com/justtaldevelops/go-hcaptcha/solver"
	"github.com/justtaldevelops/go-hcaptcha/utils"
)

// GuessSolver solves hCaptcha tasks by guessing the solution.
type GuessSolver struct{}

// Solve ...
func (s *GuessSolver) Solve(_, _ string, tasks []solver.Task) (answers []solver.Task) {
	for _, task := range tasks {
		if utils.Chance(0.5) {
			answers = append(answers, task)
		}
	}
	return answers
}
