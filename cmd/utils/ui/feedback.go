package ui

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	spinner  spinner.Model
	quitting bool
}

func (m model) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		m.quitting = true
		return m, tea.Quit
	case tea.Cmd:
		return m, msg
	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
}

func (m model) View() string {
	if m.quitting {
		return ""
	}
	str := fmt.Sprintf("\n %s Generating your resource...", m.spinner.View())
	return str
}

// ShowSpinner shows an animated spinner until the provided function completes.
func ShowSpinner(complete chan bool) {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(PrimaryColor)
	m := model{spinner: s}

	p := tea.NewProgram(m)

	go func() {
		<-complete
		p.Quit()
	}()

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running spinner:", err)
	}
}

// ActionDelay adds a small delay to make animations feel more natural
func ActionDelay() {
	time.Sleep(300 * time.Millisecond)
}
