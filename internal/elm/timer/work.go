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

func initWorkModel(interval time.Duration) WorkModel {
	return WorkModel{
		interval: interval,
		timer:    timer.New(interval),
		keymaps: workModelKeymaps{
			toggle: key.NewBinding(key.WithKeys("s"), key.WithHelp("s", "Start/Pause")),
			reset:  key.NewBinding(key.WithKeys("r"), key.WithHelp("r", "Reset")),
		},
		help: help.New(),
	}
}

type WorkModel struct {
	interval time.Duration
	timer    timer.Model
	keymaps  workModelKeymaps
	help     help.Model
}

var _ tea.Model = WorkModel{}

func (m WorkModel) Init() tea.Cmd {
	return m.timer.Init()
}

func (m WorkModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case timer.TickMsg, timer.StartStopMsg:
		var cmd tea.Cmd
		m.timer, cmd = m.timer.Update(msg)
		return m, cmd
	case timer.TimeoutMsg:
		return m, WorkCompleted
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keymaps.toggle):
			return m, m.timer.Toggle()
		case key.Matches(msg, m.keymaps.reset):
			return m.Reset(), nil
		}
	}
	return m, nil
}

func (m WorkModel) View() string {
	title := lipgloss.NewStyle().Bold(true).SetString("🍅 Pomodoro Timer")
	return fmt.Sprintf("%s\n", title) +
		"\n" +
		fmt.Sprintf("Working on it... %s Remaining\n", m.timer.View()) +
		"\n" +
		m.help.ShortHelpView(m.keymaps.bindings())
}

func (m WorkModel) Reset() WorkModel {
	m.timer.Timeout = m.interval
	return m
}

type workModelKeymaps struct {
	toggle key.Binding
	reset  key.Binding
}

func (k workModelKeymaps) bindings() []key.Binding {
	return []key.Binding{k.toggle, k.reset}
}
