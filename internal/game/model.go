package game

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"

	"integra/pkg/quiz"
)

type screenMode int

const (
	screenMenu screenMode = iota
	screenQuiz
	screenError
)

type Model struct {
	quizzes      []quiz.Quiz
	selectedQuiz int
	currentQuiz  int
	questions    []quiz.Question
	current      int
	selected     int
	score        int
	showAnswer   bool
	lastCorrect  bool
	finished     bool
	width        int
	height       int
	theme        theme
	screen       screenMode
	loadError    string
}

func New() Model {
	model := Model{
		theme:       newTheme(),
		currentQuiz: -1,
		screen:      screenMenu,
	}
	return model.loadQuizzes()
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
	switch m.screen {
	case screenMenu:
		return m.handleMenuKey(key)
	case screenError:
		return m.handleErrorKey(key)
	}

	if m.finished {
		switch key {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "r":
			return m.reset(), nil
		case "m":
			return m.goToMenu(), nil
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
		case "m":
			return m.goToMenu(), nil
		default:
			return m, nil
		}
	}

	switch key {
	case "q", "ctrl+c":
		return m, tea.Quit
	case "m":
		return m.goToMenu(), nil
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

func (m Model) handleMenuKey(key string) (tea.Model, tea.Cmd) {
	switch key {
	case "q", "ctrl+c":
		return m, tea.Quit
	case "up", "k":
		if m.selectedQuiz > 0 {
			m.selectedQuiz--
		}
	case "down", "j":
		if m.selectedQuiz < len(m.quizzes)-1 {
			m.selectedQuiz++
		}
	case "1", "2", "3", "4", "5", "6", "7", "8", "9":
		index := int(key[0] - '1')
		if index >= 0 && index < len(m.quizzes) {
			m.selectedQuiz = index
			return m.startQuiz(index), nil
		}
	case "enter", " ":
		if len(m.quizzes) > 0 {
			return m.startQuiz(m.selectedQuiz), nil
		}
	case "r":
		return m.loadQuizzes(), nil
	}

	return m, nil
}

func (m Model) handleErrorKey(key string) (tea.Model, tea.Cmd) {
	switch key {
	case "q", "ctrl+c":
		return m, tea.Quit
	case "r":
		return m.loadQuizzes(), nil
	}
	return m, nil
}

func (m Model) currentQuestion() quiz.Question {
	return m.questions[m.current]
}

func (m Model) currentQuizData() quiz.Quiz {
	if m.currentQuiz >= 0 && m.currentQuiz < len(m.quizzes) {
		return m.quizzes[m.currentQuiz]
	}
	return quiz.Quiz{}
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
	if m.currentQuiz >= 0 && m.currentQuiz < len(m.quizzes) {
		return m.startQuiz(m.currentQuiz)
	}
	return m.goToMenu()
}

func (m Model) goToMenu() Model {
	m.screen = screenMenu
	m.finished = false
	m.showAnswer = false
	m.selected = 0
	m.current = 0
	if m.currentQuiz >= 0 {
		m.selectedQuiz = m.currentQuiz
	}
	return m
}

func (m Model) startQuiz(index int) Model {
	if index < 0 || index >= len(m.quizzes) {
		return m
	}

	selectedQuiz := m.quizzes[index]
	m.currentQuiz = index
	m.selectedQuiz = index
	m.questions = append([]quiz.Question(nil), selectedQuiz.Questions...)
	m.current = 0
	m.selected = 0
	m.score = 0
	m.showAnswer = false
	m.lastCorrect = false
	m.finished = false
	m.screen = screenQuiz
	m.loadError = ""
	return m
}

func (m Model) loadQuizzes() Model {
	quizzes, err := quiz.LoadDir("questions")
	if err != nil {
		m.quizzes = nil
		m.questions = nil
		m.currentQuiz = -1
		m.screen = screenError
		m.loadError = fmt.Sprintf("Erro ao carregar quizzes: %v", err)
		return m
	}

	m.quizzes = quizzes
	m.questions = nil
	m.currentQuiz = -1
	m.selectedQuiz = 0
	m.screen = screenMenu
	m.finished = false
	m.showAnswer = false
	m.current = 0
	m.selected = 0
	m.score = 0
	m.loadError = ""
	return m
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
