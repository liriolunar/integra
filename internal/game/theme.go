package game

import "github.com/charmbracelet/lipgloss"

// Solarized accent colors (consistent across light and dark modes).
var (
	solYellow  = lipgloss.Color("#b58900")
	solOrange  = lipgloss.Color("#cb4b16")
	solRed     = lipgloss.Color("#dc322f")
	solMagenta = lipgloss.Color("#d33682")
	solViolet  = lipgloss.Color("#6c71c4")
	solBlue    = lipgloss.Color("#268bd2")
	solCyan    = lipgloss.Color("#2aa198")
	solGreen   = lipgloss.Color("#859900")
)

// Adaptive Solarized base colors that switch between light and dark
// terminal backgrounds.
var (
	solBaseFg          = lipgloss.AdaptiveColor{Light: "#657b83", Dark: "#839496"} // Base00 / Base0
	solBaseEmphasis    = lipgloss.AdaptiveColor{Light: "#586e75", Dark: "#93a1a1"} // Base01 / Base1
	solBaseSecondary   = lipgloss.AdaptiveColor{Light: "#93a1a1", Dark: "#586e75"} // Base1 / Base01
	solBaseInverse     = lipgloss.AdaptiveColor{Light: "#002b36", Dark: "#fdf6e3"} // Base03 / Base3
	solBaseBgHighlight = lipgloss.AdaptiveColor{Light: "#eee8d5", Dark: "#073642"} // Base2 / Base02
)

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
			Foreground(solBaseFg),
		title: lipgloss.NewStyle().
			Bold(true).
			Foreground(solBaseEmphasis).
			Padding(0, 2),
		subtitle: lipgloss.NewStyle().
			Bold(true).
			Foreground(solViolet),
		stat: lipgloss.NewStyle().
			Bold(true).
			Foreground(solCyan).
			Padding(0, 1).
			MarginRight(1),
		topic: lipgloss.NewStyle().
			Bold(true).
			Foreground(solYellow).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(solOrange).
			Padding(0, 2),
		questionBox: lipgloss.NewStyle().
			Bold(true).
			Foreground(solBaseEmphasis).
			Border(lipgloss.DoubleBorder()).
			BorderForeground(solBlue).
			Padding(1, 2).
			MarginTop(1).
			MarginBottom(1),
		info: lipgloss.NewStyle().
			Foreground(solBaseSecondary).
			MarginTop(1),
		success: lipgloss.NewStyle().
			Bold(true).
			Foreground(solGreen),
		error: lipgloss.NewStyle().
			Bold(true).
			Foreground(solRed),
		explanation: lipgloss.NewStyle().
			Foreground(solBaseFg).
			Border(lipgloss.NormalBorder()).
			BorderForeground(solViolet).
			Padding(1, 2).
			MarginTop(1),
		gap: lipgloss.NewStyle(),
		menuItem: lipgloss.NewStyle().
			Foreground(solBaseFg).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(solBaseSecondary).
			Padding(1, 2).
			MarginBottom(1),
		menuSelect: lipgloss.NewStyle().
			Foreground(solBaseInverse).
			Background(solYellow).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(solYellow).
			Padding(1, 2).
			MarginBottom(1).
			Bold(true),
		muted: lipgloss.NewStyle().
			Foreground(solBaseSecondary),
	}
}

func (t theme) choiceDefault() lipgloss.Style {
	return t.choiceBase().
		Foreground(solBaseFg).
		BorderForeground(solBaseSecondary)
}

func (t theme) choiceSelected() lipgloss.Style {
	return t.choiceBase().
		Foreground(solBaseInverse).
		Background(solYellow).
		BorderForeground(solYellow)
}

func (t theme) choiceCorrect() lipgloss.Style {
	return t.choiceBase().
		Foreground(solBaseInverse).
		Background(solGreen).
		BorderForeground(solGreen)
}

func (t theme) choiceWrong() lipgloss.Style {
	return t.choiceBase().
		Foreground(solBaseInverse).
		Background(solRed).
		BorderForeground(solRed)
}

func (t theme) choiceBase() lipgloss.Style {
	return lipgloss.NewStyle().
		Padding(1, 1).
		Align(lipgloss.Center, lipgloss.Center).
		Bold(true).
		Border(lipgloss.RoundedBorder())
}
