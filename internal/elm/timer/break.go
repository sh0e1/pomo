package timer

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func initBreakModel(interval time.Duration) BreakModel {
	return BreakModel{
		timer: timer.New(interval),
		keymaps: breakModelKeymaps{
			start: key.NewBinding(key.WithKeys("b"), key.WithHelp("b", "Start")),
			skip:  key.NewBinding(key.WithKeys("s"), key.WithHelp("s", "Skip")),
		},
		help: help.New(),
	}
}

type BreakModel struct {
	timer   timer.Model
	keymaps breakModelKeymaps
	help    help.Model
}

var _ tea.Model = BreakModel{}

func (m BreakModel) Init() tea.Cmd {
	return nil
}

func (m BreakModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case timer.TickMsg, timer.StartStopMsg:
		var cmd tea.Cmd
		m.timer, cmd = m.timer.Update(msg)
		return m, cmd
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keymaps.start):
			return m, m.timer.Start()
		}
	}
	return m, nil
}

func (m BreakModel) View() string {
	title := lipgloss.NewStyle().Bold(true).SetString("☕️ Break Time")
	return fmt.Sprintf("%s\n", title) +
		"\n" +
		fmt.Sprintf("Take a break... %s Remaining\n", m.timer.View()) +
		"\n" +
		m.help.ShortHelpView(m.keymaps.bindings())
}

type breakModelKeymaps struct {
	start key.Binding
	skip  key.Binding
}

func (k breakModelKeymaps) bindings() []key.Binding {
	return []key.Binding{k.start, k.skip}
}
