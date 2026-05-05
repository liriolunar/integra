package game

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"

	"integra/pkg/quiz"
)

func (m Model) View() string {
	var content string
	switch m.screen {
	case screenMenu:
		content = m.menuView()
	case screenError:
		content = m.errorView()
	default:
		content = m.finishView()
		if !m.finished {
			content = m.gameView()
		}
	}

	rendered := outerWidth(m.theme.app, m.contentWidth()+m.theme.app.GetHorizontalFrameSize()).
		Render(content)
	if m.width == 0 || m.height == 0 {
		return rendered
	}

	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Top, rendered)
}

func (m Model) gameView() string {
	q := m.currentQuestion()
	subtitle := "limites • derivadas • integrais"
	if currentQuiz := m.currentQuizData(); currentQuiz.Name != "" {
		subtitle = currentQuiz.Name
	}
	sections := []string{
		m.theme.title.Render("INTEGRA"),
		m.theme.subtitle.Render(subtitle),
		"",
		m.statsView(q),
		m.questionView(q),
		m.choicesView(q),
		m.footerView(q),
	}

	return strings.Join(sections, "\n")
}

func (m Model) menuView() string {
	items := make([]string, 0, len(m.quizzes))
	for index, loadedQuiz := range m.quizzes {
		style := m.theme.menuItem
		if index == m.selectedQuiz {
			style = m.theme.menuSelect
		}

		content := fmt.Sprintf("%d. %s", index+1, loadedQuiz.Name)
		if loadedQuiz.Description != "" {
			content += "\n" + loadedQuiz.Description
		}
		content += fmt.Sprintf("\n%d perguntas", len(loadedQuiz.Questions))
		items = append(items, outerWidth(style, m.contentWidth()).Render(content))
	}

	sections := []string{
		m.theme.title.Render("INTEGRA"),
		m.theme.subtitle.Render("Selecione um quiz em /questions"),
		"",
		strings.Join(items, "\n"),
		m.theme.info.Render("Use ↑/↓ ou j/k para navegar • Enter para abrir • 1-9 para atalho • q para sair"),
	}

	return strings.Join(sections, "\n")
}

func (m Model) errorView() string {
	card := outerWidth(m.theme.explanation, min(70, m.contentWidth())).Render(m.loadError)
	return strings.Join([]string{
		m.theme.title.Render("INTEGRA"),
		m.theme.subtitle.Render("Erro ao carregar quizzes"),
		"",
		card,
		m.theme.info.Render("Pressione r para tentar recarregar ou q para sair."),
	}, "\n")
}

func (m Model) statsView(q quiz.Question) string {
	stats := []string{
		m.theme.stat.Render(fmt.Sprintf("PERGUNTA %d/%d", m.current+1, len(m.questions))),
		m.theme.stat.Render(fmt.Sprintf("PONTOS %d", m.score)),
	}

	rows := []string{
		m.joinHorizontalWithGap(1, stats...),
		m.theme.topic.Render(strings.ToUpper(q.Topic)),
	}

	return strings.Join(rows, "\n")
}

func (m Model) questionView(q quiz.Question) string {
	return outerWidth(m.theme.questionBox, m.contentWidth()).
		Render("QUESTÃO\n\n" + q.Prompt)
}

func (m Model) choicesView(q quiz.Question) string {
	columns := m.choiceColumns()
	width := m.choiceWidth(columns)
	rows := make([]string, 0, (len(q.Options)+columns-1)/columns)

	for start := 0; start < len(q.Options); start += columns {
		end := min(start+columns, len(q.Options))
		rowCards := make([]string, 0, end-start)

		for index := start; index < end; index++ {
			rowCards = append(rowCards, m.renderChoice(q, index, width))
		}

		rows = append(rows, m.joinHorizontalWithGap(2, rowCards...))
	}

	return strings.Join(rows, "\n")
}

func (m Model) renderChoice(q quiz.Question, index, width int) string {
	style := m.theme.choiceDefault()
	switch {
	case m.showAnswer && index == q.Answer:
		style = m.theme.choiceCorrect()
	case m.showAnswer && index == m.selected && !m.lastCorrect:
		style = m.theme.choiceWrong()
	case index == m.selected:
		style = m.theme.choiceSelected()
	}

	height := 7
	if width < 22 {
		height = 8
	}

	style = outerWidth(style.Height(height), width)

	letter := string(rune('A' + index))
	body := fmt.Sprintf("%s\n\n%s", letter, q.Options[index])

	return style.Render(body)
}

func (m Model) footerView(q quiz.Question) string {
	if !m.showAnswer {
		controls := "Use ←/→ ou h/l para mover • a/b/c/d para escolher • Enter/Espaço para confirmar • m menu • q sair"
		if m.choiceColumns() == 1 {
			controls = "Use h/l ou a/b/c/d para mudar a opção • Enter/Espaço para confirmar • m menu • q sair"
		}
		return m.theme.info.Render(controls)
	}

	result := m.theme.error.Render("✗ Resposta incorreta")
	if m.lastCorrect {
		result = m.theme.success.Render("✓ Resposta correta")
	}

	explanation := outerWidth(m.theme.explanation, m.contentWidth()).
		Render(fmt.Sprintf("Resposta certa: %s) %s\n\n%s", answerLetter(q.Answer), q.Options[q.Answer], q.Explanation))

	return strings.Join([]string{
		result,
		explanation,
		m.theme.info.Render("Pressione Enter, Espaço, → ou l para continuar • m para voltar ao menu."),
	}, "\n")
}

func (m Model) finishView() string {
	total := len(m.questions)
	percent := float64(m.score) / float64(total) * 100

	rating := "Bom começo — continue praticando integrais e limites."
	switch {
	case percent >= 90:
		rating = "Excelente — você dominou o modo difícil."
	case percent >= 70:
		rating = "Muito bom — seu cálculo está afiado."
	case percent >= 50:
		rating = "Boa base — mais algumas rodadas e você sobe de nível."
	}

	summaryCard := outerWidth(m.theme.questionBox, min(60, m.contentWidth())).
		Render(rating)

	return strings.Join([]string{
		m.theme.title.Render("INTEGRA"),
		m.theme.subtitle.Render(m.currentQuizData().Name),
		"",
		m.theme.stat.Render(fmt.Sprintf("PONTOS %d/%d", m.score, total)),
		m.theme.stat.Render(fmt.Sprintf("ACERTO %.0f%%", percent)),
		"",
		summaryCard,
		m.theme.info.Render("Pressione r para jogar novamente • m para menu • q para sair."),
	}, "\n")
}

func answerLetter(index int) string {
	return string(rune('a' + index))
}

func outerWidth(style lipgloss.Style, totalWidth int) lipgloss.Style {
	innerWidth := totalWidth - style.GetHorizontalFrameSize()
	if innerWidth < 1 {
		innerWidth = 1
	}
	return style.Width(innerWidth)
}

func (m Model) joinHorizontalWithGap(gapWidth int, parts ...string) string {
	if len(parts) == 0 {
		return ""
	}
	if len(parts) == 1 {
		return parts[0]
	}

	items := make([]string, 0, len(parts)*2-1)
	gap := m.theme.gap.Width(gapWidth).Render("")
	for index, part := range parts {
		if index > 0 {
			items = append(items, gap)
		}
		items = append(items, part)
	}

	return lipgloss.JoinHorizontal(lipgloss.Top, items...)
}
