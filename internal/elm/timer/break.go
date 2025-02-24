package timer

import (
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
)

func initBreakModel(cfg *Config) BreakModel {
	return BreakModel{
		shortInterval: cfg.ShortBreakInterval,
		longInterval:  cfg.LongBreakInterval,
		timer:         timer.New(cfg.ShortBreakInterval),
		keymaps: breakModelKeymaps{
			start: key.NewBinding(key.WithKeys("b"), key.WithHelp("b", "Start Break")),
			skip:  key.NewBinding(key.WithKeys("s"), key.WithHelp("s", "Skip Break")),
		},
		help: help.New(),
	}
}

type BreakModel struct {
	shortInterval time.Duration
	longInterval  time.Duration
	timer         timer.Model
	keymaps       breakModelKeymaps
	help          help.Model
}

var _ tea.Model = BreakModel{}

func (m BreakModel) Init() tea.Cmd {
	return m.timer.Init()
}

func (m BreakModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m BreakModel) View() string {
	return "break time"
}

type breakModelKeymaps struct {
	start key.Binding
	skip  key.Binding
}
