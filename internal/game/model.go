package game

import (
	tea "github.com/charmbracelet/bubbletea"

	"proj-aplicado-6/pkg/quiz"
)

type Model struct {
	questions   []quiz.Question
	current     int
	selected    int
	score       int
	showAnswer  bool
	lastCorrect bool
	finished    bool
	width       int
	height      int
	theme       theme
}

func New() Model {
	questions := append([]quiz.Question(nil), quiz.Bank...)

	return Model{
		questions: questions,
		theme:     newTheme(),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil
	case tea.KeyMsg:
		return m.handleKey(msg.String())
	default:
		return m, nil
	}
}

func (m Model) handleKey(key string) (tea.Model, tea.Cmd) {
	if m.finished {
		switch key {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "r":
			return m.reset(), nil
		default:
			return m, nil
		}
	}

	if m.showAnswer {
		switch key {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter", " ", "right", "l":
			return m.advanceQuestion(), nil
		default:
			return m, nil
		}
	}

	switch key {
	case "q", "ctrl+c":
		return m, tea.Quit
	case "left", "h":
		if m.selected > 0 {
			m.selected--
		}
	case "right", "l":
		if m.selected < len(m.currentQuestion().Options)-1 {
			m.selected++
		}
	case "a", "b", "c", "d":
		m.selectByLetter(key)
	case "enter", " ":
		m.lastCorrect = m.selected == m.currentQuestion().Answer
		if m.lastCorrect {
			m.score++
		}
		m.showAnswer = true
	}

	return m, nil
}

func (m Model) currentQuestion() quiz.Question {
	return m.questions[m.current]
}

func (m *Model) selectByLetter(key string) {
	index := int(key[0] - 'a')
	if index >= 0 && index < len(m.currentQuestion().Options) {
		m.selected = index
	}
}

func (m Model) advanceQuestion() Model {
	m.showAnswer = false
	m.selected = 0
	m.current++
	if m.current >= len(m.questions) {
		m.finished = true
	}
	return m
}

func (m Model) reset() Model {
	next := New()
	next.width = m.width
	next.height = m.height
	return next
}

func (m Model) contentWidth() int {
	if m.width == 0 {
		return 96
	}

	contentWidth := m.width - 6
	if contentWidth < 1 {
		return 1
	}
	if contentWidth > 96 {
		return 96
	}
	return contentWidth
}

func (m Model) choiceColumns() int {
	switch {
	case m.width >= 128:
		return 4
	case m.width >= 76:
		return 2
	default:
		return 1
	}
}

func (m Model) choiceWidth(columns int) int {
	availableWidth := m.contentWidth()
	gapWidth := 2 * (columns - 1)
	cardWidth := (availableWidth - gapWidth) / columns
	if cardWidth < 1 {
		return 1
	}
	return cardWidth
}
