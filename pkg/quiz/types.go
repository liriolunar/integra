package quiz

type Question struct {
	Topic       string   `json:"topic" yaml:"topic"`
	Prompt      string   `json:"prompt" yaml:"prompt"`
	Options     []string `json:"options" yaml:"options"`
	Answer      int      `json:"answer" yaml:"answer"`
	Explanation string   `json:"explanation" yaml:"explanation"`
}

type Quiz struct {
	Name        string     `json:"name" yaml:"name"`
	Description string     `json:"description" yaml:"description"`
	Questions   []Question `json:"questions" yaml:"questions"`
}
