package solver

// Solver is an interface to solve hCaptcha tasks.
type Solver interface {
	// Solve solves the hCaptcha tasks using the category, question, and the task. If it was successful,
	// it returns true, and in all other cases, it returns false.
	Solve(category, question string, tasks []Task) []Task
}

// Task is a task assigned by hCaptcha.
type Task struct {
	// Image is the image to represent the task.
	Image []byte
	// Key is the task key, used when referencing answers.
	Key string
	// Index is the index of the task.
	Index int
}
