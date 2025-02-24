package timer

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
)

func initWorkModel(cfg *Config) tea.Model {
	return WorkModel{
		interval: cfg.WorkInterval,
		timer:    timer.New(cfg.WorkInterval),
		keymaps: keymaps{
			toggle: key.NewBinding(key.WithKeys("s"), key.WithHelp("s", "Start/Pause")),
			reset:  key.NewBinding(key.WithKeys("r"), key.WithHelp("r", "Reset")),
			quit:   key.NewBinding(key.WithKeys("q"), key.WithHelp("q", "Quit")),
		},
		help: help.New(),
	}
}

type WorkModel struct {
	interval time.Duration
	timer    timer.Model
	keymaps  keymaps
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
		return m, tea.Quit
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keymaps.toggle):
			return m, m.timer.Toggle()
		case key.Matches(msg, m.keymaps.reset):
			m.timer.Timeout = m.interval
			return m, nil
		case key.Matches(msg, m.keymaps.quit):
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m WorkModel) View() string {
	return fmt.Sprintf("Working on it... %s Remaining\n%s", m.timer.View(), m.help.ShortHelpView(m.keymaps.bindings()))
}

type keymaps struct {
	toggle key.Binding
	reset  key.Binding
	quit   key.Binding
}

func (k keymaps) bindings() []key.Binding {
	return []key.Binding{k.toggle, k.reset, k.quit}
}
