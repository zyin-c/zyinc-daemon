package runner

type RunnerInstances struct {
	BaseDirectory string `json:"basedir"`
	FileName      string `json:"filename"`
	EnvFile       string `json:"envfile"`
}

type Runner struct {
	RunnerId        string
	RunnerInstances map[string]RunnerInstances
}

func NewRunner() (*Runner, error) {
	return &Runner{}, nil
}
