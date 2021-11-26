package manual

import (
	"github.com/justtaldevelops/go-hcaptcha/solver"
)

// ManualSolver solves hCaptcha tasks by manual.
type ManualSolver struct{
	SolveFunc func (category, object string, tasks []solver.Task) (answers []solver.Task)
}

// Solve ...
func (s *ManualSolver) Solve(category, object string, tasks []solver.Task) (answers []solver.Task) {
	return  s.SolveFunc(category, object, tasks)
}
