package game

import "github.com/charmbracelet/lipgloss"

type theme struct {
	app         lipgloss.Style
	title       lipgloss.Style
	subtitle    lipgloss.Style
	stat        lipgloss.Style
	topic       lipgloss.Style
	questionBox lipgloss.Style
	info        lipgloss.Style
	success     lipgloss.Style
	error       lipgloss.Style
	explanation lipgloss.Style
	gap         lipgloss.Style
	menuItem    lipgloss.Style
	menuSelect  lipgloss.Style
	muted       lipgloss.Style
}

func newTheme() theme {
	return theme{
		app: lipgloss.NewStyle().
			Padding(1, 2).
			Foreground(lipgloss.Color("#F8FAFC")),
		title: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#F8FAFC")).
			Padding(0, 2),
		subtitle: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#C4B5FD")),
		stat: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#0F172A")).
			Foreground(lipgloss.Color("#67E8F9")).
			Padding(0, 1).
			MarginRight(1),
		topic: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FDE68A")).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#F59E0B")).
			Padding(0, 2),
		questionBox: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#F8FAFC")).
			Border(lipgloss.DoubleBorder()).
			BorderForeground(lipgloss.Color("#38BDF8")).
			Padding(1, 2).
			MarginTop(1).
			MarginBottom(1),
		info: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#CBD5E1")).
			MarginTop(1),
		success: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#86EFAC")),
		error: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FDA4AF")),
		explanation: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#E2E8F0")).
			Border(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("#7C3AED")).
			Padding(1, 2).
			MarginTop(1),
		gap: lipgloss.NewStyle(),
		menuItem: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#E2E8F0")).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#475569")).
			Padding(1, 2).
			MarginBottom(1),
		menuSelect: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#0F172A")).
			Background(lipgloss.Color("#FDE047")).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#FACC15")).
			Padding(1, 2).
			MarginBottom(1).
			Bold(true),
		muted: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#94A3B8")),
	}
}

func (t theme) choiceDefault() lipgloss.Style {
	return t.choiceBase().
		Foreground(lipgloss.Color("#E2E8F0")).
		BorderForeground(lipgloss.Color("#475569"))
}

func (t theme) choiceSelected() lipgloss.Style {
	return t.choiceBase().
		Foreground(lipgloss.Color("#0F172A")).
		Background(lipgloss.Color("#FDE047")).
		BorderForeground(lipgloss.Color("#FACC15"))
}

func (t theme) choiceCorrect() lipgloss.Style {
	return t.choiceBase().
		Foreground(lipgloss.Color("#ECFDF5")).
		Background(lipgloss.Color("#166534")).
		BorderForeground(lipgloss.Color("#4ADE80"))
}

func (t theme) choiceWrong() lipgloss.Style {
	return t.choiceBase().
		Foreground(lipgloss.Color("#FEF2F2")).
		Background(lipgloss.Color("#991B1B")).
		BorderForeground(lipgloss.Color("#F87171"))
}

func (t theme) choiceBase() lipgloss.Style {
	return lipgloss.NewStyle().
		Padding(1, 1).
		Align(lipgloss.Center, lipgloss.Center).
		Bold(true).
		Border(lipgloss.RoundedBorder())
}
